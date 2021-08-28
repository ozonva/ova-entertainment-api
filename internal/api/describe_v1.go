package api

import (
	"context"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"time"
)

func (s *ApiServer) DescribeEntertainmentV1(ctx context.Context, req *desc.DescribeEntertainmentV1Request) (*desc.EntertainmentV1Response, error) {

	entertainment, err := s.repo.DescribeEntertainment(models.Entertainment{
		ID:          req.ID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Date:        time.Time{},
	})
	if err != nil {
		return nil, err
	}

	return entertainment.Marshall(), nil
}
