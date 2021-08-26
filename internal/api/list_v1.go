package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) ListEntertainmentsV1(ctx context.Context, req *desc.ListEntertainmentV1Request) (*desc.ListEntertainmentsV1Response, error) {

	log.Info().Msg("Run ListEntertainmentsV1")

	return &desc.ListEntertainmentsV1Response{
		List: []*desc.EntertainmentV1Response{},
	}, nil
}
