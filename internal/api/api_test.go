package api

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Api", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *repo.MockRepo
		cntx     context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRepo(mockCtrl)
		cntx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("API: CreateEntertainmentV1", func() {
		It("should not error", func() {

			mockRepo.EXPECT().AddEntertainments([]models.Entertainment{
				models.New(2, "Title", "Description"),
			}).Return(nil).Times(1)

			api := NewApiServer(mockRepo)
			_, err := api.CreateEntertainmentV1(cntx, &desc.CreateEntertainmentV1Request{
				UserID:      2,
				Title:       "Title",
				Description: "Description",
			})

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("API: DescribeEntertainmentV1", func() {
		It("should not error", func() {

			model := models.New(2, "Title", "Description")
			model.ID = 1
			mockRepo.EXPECT().DescribeEntertainment(model).Return(&model, nil).Times(1)

			api := NewApiServer(mockRepo)
			_, err := api.DescribeEntertainmentV1(cntx, &desc.DescribeEntertainmentV1Request{
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

			api := NewApiServer(mockRepo)
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

			api := NewApiServer(mockRepo)
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
