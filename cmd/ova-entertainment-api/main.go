package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-entertainment-api/internal/api"
	"github.com/ozonva/ova-entertainment-api/internal/db"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"google.golang.org/grpc"
	"log"
	"net"
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
	desc.RegisterApiServer(s, api.NewApiServer(dbConn))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {

	dsn := "postgres://db_user:db_password@localhost:5434/db?sslmode=disable"
	dbConn := db.Connect(dsn)

	if err := run(dbConn); err != nil {
		log.Fatal(err)
	}
}
