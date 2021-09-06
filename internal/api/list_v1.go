package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
)

// ListEntertainmentsV1 Получение списка сущностей
// Успешные ответы отправляются в метрику
func (s *ApiServer) ListEntertainmentsV1(ctx context.Context, req *desc.ListEntertainmentV1Request) (res *desc.ListEntertainmentsV1Response, err error) {

	defer func() {
		if err == nil {
			s.metrics.IncCounterSuccessResponseForList()
		}
	}()

	log.Info().
		Caller().
		Uint32("Limit", req.Limit).
		Uint32("Offset", req.Offset).
		Msg("")

	models, err := s.repo.ListEntertainments(req.Limit, req.Offset)
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		return nil, err
	}

	list := make([]*desc.EntertainmentV1Response, 0, req.Limit)

	for _, e := range models {
		list = append(list, e.Marshall())
	}

	return &desc.ListEntertainmentsV1Response{List: list}, nil
}
