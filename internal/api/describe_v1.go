package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) DescribeEntertainmentV1(ctx context.Context, req *desc.DescribeEntertainmentV1Request) (*desc.EntertainmentV1Response, error) {

	log.Info().Msg("Run DescribeEntertainmentV1")

	return &desc.EntertainmentV1Response{
		ID:          42,
		UserID:      1,
		Title:       "Football",
		Description: "Real Madrid vs Barcelona",
	}, nil
}
