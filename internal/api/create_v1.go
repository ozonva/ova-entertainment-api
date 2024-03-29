package api

import (
	"context"
	"github.com/ozonva/ova-entertainment-api/internal/flusher"
	"github.com/ozonva/ova-entertainment-api/internal/kafka"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	pgkserver "github.com/ozonva/ova-entertainment-api/internal/saver"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateEntertainmentV1 Добавление сущности в базу и в очередь сообщений
// При успешном добавлении отправляется метрика
func (s *ApiServer) CreateEntertainmentV1(ctx context.Context, req *desc.CreateEntertainmentV1Request) (res *emptypb.Empty, err error) {

	defer func() {
		if err == nil {
			s.metrics.IncCounterSuccessResponseForCreate()
		}
	}()

	log.Info().
		Caller().
		Uint64("UserID", req.UserID).
		Str("Title", req.Title).
		Str("Description", req.Title).
		Msg("")

	f := flusher.NewFlusher(2, s.repo)
	saver := pgkserver.NewSaver(ctx, 10, f)
	defer saver.Close()

	model := models.New(req.UserID, req.Title, req.Description)
	err = saver.Save(model)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	err = s.producer.Send(kafka.Message{
		EventType: kafka.Create,
		Value:     model,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
