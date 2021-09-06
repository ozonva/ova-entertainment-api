package api

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-entertainment-api/internal/kafka"
	"github.com/ozonva/ova-entertainment-api/internal/metrics"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/stretchr/testify/assert"
	"time"
)

var _ = Describe("Api", func() {
	var (
		mockCtrl     *gomock.Controller
		mockRepo     *repo.MockRepo
		mockProducer *kafka.MockProducer
		mockMetrics  *metrics.MockMetrics
		cntx         context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRepo(mockCtrl)
		mockProducer = kafka.NewMockProducer(mockCtrl)
		mockMetrics = metrics.NewMockMetrics(mockCtrl)
		cntx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("API: CreateEntertainmentV1", func() {
		It("should not error", func() {

			model := models.New(2, "Title", "Description")

			mockProducer.EXPECT().Send(kafka.Message{
				EventType: kafka.Create,
				Value:     model,
			}).Times(1)
			mockMetrics.EXPECT().CreateSuccessResponseIncCounter().Times(1)

			mockRepo.EXPECT().AddEntertainments([]models.Entertainment{model}).Return(nil).Times(1)

			api := NewApiServer(mockRepo, mockProducer, mockMetrics)
			_, err := api.CreateEntertainmentV1(cntx, &desc.CreateEntertainmentV1Request{
				UserID:      2,
				Title:       "Title",
				Description: "Description",
			})

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("API: UpdateEntertainmentV1", func() {
		It("should not error", func() {

			model := models.New(2, "Title", "Description")
			model.ID = 1
			model.Date = time.Now().Truncate(24*time.Hour).AddDate(0, 0, 7)

			mockRepo.EXPECT().UpdateEntertainment(model).Return(&model, nil).Times(1)

			mockProducer.EXPECT().Send(kafka.Message{
				EventType: kafka.Update,
				Value:     model.ID,
			}).Times(1)
			mockMetrics.EXPECT().UpdateSuccessResponseIncCounter().Times(1)

			api := NewApiServer(mockRepo, mockProducer, mockMetrics)
			_, err := api.UpdateEntertainmentV1(cntx, &desc.UpdateEntertainmentV1Request{
				ID:          1,
				UserID:      2,
				Title:       "Title",
				Description: "Description",
			})

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("API: RemoveEntertainmentV1", func() {
		It("should not error", func() {

			ID := uint64(1)
			mockRepo.EXPECT().RemoveEntertainment(ID).Return(nil).Times(1)

			mockProducer.EXPECT().Send(kafka.Message{
				EventType: kafka.Remove,
				Value:     ID,
			}).Times(1)
			mockMetrics.EXPECT().RemoveSuccessResponseIncCounter().Times(1)

			api := NewApiServer(mockRepo, mockProducer, mockMetrics)
			_, err := api.RemoveEntertainmentV1(cntx, &desc.RemoveEntertainmentV1Request{ID: 1})

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("API: ListEntertainmentsV1", func() {
		It("should not error", func() {

			limit := uint32(100)
			offset := uint32(10)

			models := dataProviderEntities()

			mockRepo.EXPECT().ListEntertainments(limit, offset).Return(models, nil).Times(1)
			mockMetrics.EXPECT().ListSuccessResponseIncCounter().Times(1)

			api := NewApiServer(mockRepo, mockProducer, mockMetrics)
			_, err := api.ListEntertainmentsV1(cntx, &desc.ListEntertainmentV1Request{
				Limit:  uint32(100),
				Offset: uint32(10),
			})

			assert.Nil(GinkgoT(), err)
		})
	})
})

func dataProviderEntities() []models.Entertainment {
	entertainment1 := models.New(1, "title", "description")
	entertainment2 := models.New(2, "title", "description")
	entertainment3 := models.New(3, "title", "description")
	entertainment4 := models.New(4, "title", "description")
	entertainment5 := models.New(5, "title", "description")

	slice := []models.Entertainment{
		entertainment1,
		entertainment2,
		entertainment3,
		entertainment4,
		entertainment5,
	}

	return slice
}
