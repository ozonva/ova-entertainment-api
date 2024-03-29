// Package api Имплементация методов апи для grpc
package api

import (
	"github.com/ozonva/ova-entertainment-api/internal/healthcheck"
	"github.com/ozonva/ova-entertainment-api/internal/kafka"
	"github.com/ozonva/ova-entertainment-api/internal/metrics"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
)

// ApiServer Структура API
type ApiServer struct {
	desc.UnimplementedApiServer
	repo        repo.Repo
	producer    kafka.Producer
	metrics     metrics.Metrics
	healthcheck healthcheck.Healthcheck
}

func NewApiServer(repo repo.Repo, kafka kafka.Producer, metrics metrics.Metrics, healthcheck healthcheck.Healthcheck) desc.ApiServer {
	return &ApiServer{
		repo:        repo,
		producer:    kafka,
		metrics:     metrics,
		healthcheck: healthcheck,
	}
}
