package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ApiServer) HealthCheckV1(ctx context.Context, req *emptypb.Empty) (*desc.HealthV1Response, error) {

	healthDb := "OK"
	err := s.healthcheck.HealthDB()
	if err != nil {
		healthDb = "ERROR"
	}

	healthKafka := "OK"
	err = s.healthcheck.HealthDB()
	if err != nil {
		healthKafka = "ERROR"
	}

	return &desc.HealthV1Response{
		DB:    healthDb,
		Kafka: healthKafka,
	}, nil
}
