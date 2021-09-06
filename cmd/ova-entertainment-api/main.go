package main

import (
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
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
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
)

const (
	grpcPort = ":8082"
)

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

func initMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initKafka() kafka.Producer {
	conf, err := config.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	brokerList := []string{
		fmt.Sprintf("%s:%s", conf.KafkaHost, conf.KafkaPort), //"kafka:19091",
	}
	producer, err := kafka.New(brokerList)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return producer
}

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

	return db.Connect(dsn)
}

func run(dbConn *sqlx.DB) error {

	// init opentracing and jaeger
	tracer, closer := initJaeger()
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	// init producer kafka
	producer := initKafka()

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

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {

	if err := run(initDB()); err != nil {
		log.Fatal(err)
	}
}
