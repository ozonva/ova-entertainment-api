package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ApiServer) RemoveEntertainmentV1(ctx context.Context, req *desc.RemoveEntertainmentV1Request) (*emptypb.Empty, error) {

	err := s.repo.RemoveEntertainment(req.ID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
