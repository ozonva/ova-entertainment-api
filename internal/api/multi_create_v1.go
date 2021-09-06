package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-entertainment-api/internal/kafka"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/utils"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"google.golang.org/protobuf/types/known/emptypb"
)

// bulksSize на сколько батчей делить исходный спискок
const bulksSize = uint(3)

// MultiCreateEntertainmentV1 Множественное создание сущностей
// Сущности добавляются в базу батчами по bulksSize
// Батчи добавляются в трейсинг
// Успешный результат отправляется в метрику
func (s *ApiServer) MultiCreateEntertainmentV1(ctx context.Context, req *desc.MultiCreateEntertainmentV1Request) (res *emptypb.Empty, err error) {

	defer func() {
		if err == nil {
			s.metrics.IncCounterSuccessResponseForMultiCreate()
		}
	}()

	entertainments := make([]models.Entertainment, 0, len(req.Models))
	for _, model := range req.Models {
		entertainments = append(entertainments, models.New(model.UserID, model.Title, model.Description))
	}

	tracer := opentracing.GlobalTracer()
	parentSpan := tracer.StartSpan(
		"MultiCreate",
		opentracing.Tag{Key: "TotalCount", Value: len(entertainments)},
	)
	defer parentSpan.Finish()

	bulks := utils.SplitToBulks(entertainments, bulksSize)
	for _, slice := range bulks {
		err = s.repo.AddEntertainments(slice)
		if err != nil {
			return nil, err
		}

		err = s.producer.Send(kafka.Message{
			EventType: kafka.MultiCreate,
			Value:     slice,
		})
		if err != nil {
			return nil, err
		}

		span := opentracing.StartSpan(
			"MultiCreate",
			opentracing.Tag{Key: "BulkCount", Value: len(slice)},
			opentracing.ChildOf(parentSpan.Context()),
		)
		span.Finish()
	}

	return &emptypb.Empty{}, nil
}
