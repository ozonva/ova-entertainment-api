package api

import (
	"context"
	"github.com/ozonva/ova-entertainment-api/internal/kafka"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ApiServer) RemoveEntertainmentV1(ctx context.Context, req *desc.RemoveEntertainmentV1Request) (*emptypb.Empty, error) {

	defer s.metrics.RemoveSuccessResponseIncCounter()

	log.Info().
		Caller().
		Uint64("UserID", req.ID).
		Msg("")

	err := s.repo.RemoveEntertainment(req.ID)
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		return nil, err
	}

	err = s.producer.Send(kafka.Message{
		EventType: kafka.Remove,
		Value:     req.ID,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
