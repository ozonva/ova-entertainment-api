package api

import (
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
)

type ApiServer struct {
	desc.UnimplementedApiServer
	repo repo.Repo
}

func NewApiServer(r repo.Repo) desc.ApiServer {
	return &ApiServer{repo: r}
}
