package api

import desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"

type Server struct {
	desc.UnimplementedApiServer
}

func NewApiServer() desc.ApiServer {
	return &Server{}
}
