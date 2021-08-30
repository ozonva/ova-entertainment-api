package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ApiServer) RemoveEntertainmentV1(ctx context.Context, req *desc.RemoveEntertainmentV1Request) (*emptypb.Empty, error) {

	log.Info().
		Caller().
		Uint64("UserID", req.ID).
		Msg("")

	err := s.repo.RemoveEntertainment(req.ID)
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
