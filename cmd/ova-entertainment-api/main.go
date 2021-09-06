// Микросервис ova-entertainment-api
// Для общения используется GRPC
// CUD операции кладутся в kafka
// Метрики собираются в grafana и трейсинг в jaeger
package main

import (
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/halink0803/zerolog-graylog-hook/graylog"
	"github.com/jmoiron/sqlx"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-entertainment-api/internal/api"
	"github.com/ozonva/ova-entertainment-api/internal/config"
	"github.com/ozonva/ova-entertainment-api/internal/db"
	"github.com/ozonva/ova-entertainment-api/internal/kafka"
	"github.com/ozonva/ova-entertainment-api/internal/metrics"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	zerologger "github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Порт для общения по GRPC
const (
	grpcPort = ":8082"
)

// Инициализация трейсинга
func initJaeger() (opentracing.Tracer, io.Closer) {
	cfg := jaegercfg.Configuration{
		ServiceName: "ova-entertainment-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := jaegermetrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Fatalf("Can not create tracer, %s", err)
	}
	return tracer, closer
}

// Инициализация сборщика метрик
func initMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Инициализация очереди сообщений
func initKafka() kafka.Producer {
	conf, err := config.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	brokerList := []string{
		fmt.Sprintf("%s:%s", conf.KafkaHost, conf.KafkaPort),
	}
	producer, err := kafka.New(brokerList)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return producer
}

// Инициализация базы данных
func initDB() *sqlx.DB {
	conf, err := config.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)

	db, err := db.Connect(dsn)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return db
}

// Запуск GRPC сервера и инициализация сопутствующих сервисов
func run(dbConn *sqlx.DB, sigc chan os.Signal) error {

	defer dbConn.Close()

	// init opentracing and jaeger
	tracer, closer := initJaeger()
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	// init producer kafka
	producer := initKafka()
	defer producer.Close()

	// init grpc and api
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))
	desc.RegisterApiServer(s, api.NewApiServer(repo.NewRepo(dbConn), producer, metrics.NewMetrics()))
	grpc_prometheus.Register(s)

	// init prometheus metrics
	go initMetrics()

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := errgroup.Group{}
	g.Go(func() error {
		zerologger.Log().Caller().Msg("Start Server")
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	<-sigc
	zerologger.Log().Caller().Msg("Graceful shutdown")
	s.GracefulStop()

	return nil
}

func main() {

	zerologger.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	hook, err := graylog.NewGraylogHook("udp://127.0.0.1:12201")
	if err != nil {
		log.Fatal(err)
	}
	zerologger.Logger = zerologger.Hook(hook)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	if err = run(initDB(), sigc); err != nil {
		log.Fatal(err)
	}
}
