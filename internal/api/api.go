package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
)

type ApiServer struct {
	desc.UnimplementedApiServer
	repo repo.Repo
}

func NewApiServer(db *sqlx.DB) desc.ApiServer {
	r := repo.NewRepo(db)
	return &ApiServer{repo: r}
}
