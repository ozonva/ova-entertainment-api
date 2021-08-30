package api

import (
	"context"
	"github.com/ozonva/ova-entertainment-api/internal/flusher"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	pgkserver "github.com/ozonva/ova-entertainment-api/internal/saver"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ApiServer) CreateEntertainmentV1(ctx context.Context, req *desc.CreateEntertainmentV1Request) (*emptypb.Empty, error) {

	log.Info().
		Caller().
		Uint64("UserID", req.UserID).
		Str("Title", req.Title).
		Str("Description", req.Title).
		Msg("")

	f := flusher.NewFlusher(2, s.repo)
	saver := pgkserver.NewSaver(10, f)
	defer saver.Close()

	err := saver.Save(models.New(req.UserID, req.Title, req.Description))
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
