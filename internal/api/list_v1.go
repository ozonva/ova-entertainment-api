package api

import (
	"context"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
)

func (s *ApiServer) ListEntertainmentsV1(ctx context.Context, req *desc.ListEntertainmentV1Request) (*desc.ListEntertainmentsV1Response, error) {

	models, err := s.repo.ListEntertainments(req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	list := make([]*desc.EntertainmentV1Response, 0, req.Limit)

	for _, e := range models {
		list = append(list, e.Marshall())
	}

	return &desc.ListEntertainmentsV1Response{List: list}, nil
}
