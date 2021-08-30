package api

import (
	"context"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *ApiServer) DescribeEntertainmentV1(ctx context.Context, req *desc.DescribeEntertainmentV1Request) (*desc.EntertainmentV1Response, error) {

	log.Info().
		Caller().
		Uint64("ID", req.ID).
		Uint64("UserID", req.UserID).
		Str("Title", req.Title).
		Str("Description", req.Title).
		Msg("")

	entertainment, err := s.repo.DescribeEntertainment(models.Entertainment{
		ID:          req.ID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Date:        time.Now().Truncate(24*time.Hour).AddDate(0, 0, 7),
	})
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		return nil, err
	}

	return entertainment.Marshall(), nil
}
