package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ozonva/ova-entertainment-api/internal/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
)

const (
	grpcPort = ":8082"
)

func ReadCsvFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	reader := csv.NewReader(file)
	for {
		value, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(value)
	}
}

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterApiServer(s, api.NewApiServer())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
