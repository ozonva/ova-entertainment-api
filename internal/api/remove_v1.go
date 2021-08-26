package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) RemoveEntertainmentV1(ctx context.Context, req *desc.RemoveEntertainmentV1Request) (*desc.RemoveV1Response, error) {

	log.Info().Msg("Run RemoveEntertainmentV1")

	return &desc.RemoveV1Response{
		Status: true,
	}, nil
}
