package api

import (
	"context"
	"github.com/ozonva/ova-entertainment-api/internal/flusher"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ApiServer) CreateEntertainmentV1(ctx context.Context, req *desc.CreateEntertainmentV1Request) (*emptypb.Empty, error) {

	f := flusher.NewFlusher(2, s.repo)
	err := f.Flush([]models.Entertainment{
		models.New(req.UserID, req.Title, req.Description),
	})

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
