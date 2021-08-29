package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/ozonva/ova-entertainment-api/internal/api"
	"github.com/ozonva/ova-entertainment-api/internal/db"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	grpcPort = ":8082"
)

func run(dbConn *sqlx.DB) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterApiServer(s, api.NewApiServer(repo.NewRepo(dbConn)))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed open env: %v", err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@localhost:5434/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	dbConn := db.Connect(dsn)

	if err := run(dbConn); err != nil {
		log.Fatal(err)
	}
}
