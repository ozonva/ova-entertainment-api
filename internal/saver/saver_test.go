package saver

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-entertainment-api/internal/flusher"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *repo.MockRepo
		f        flusher.Flusher
		cntx     context.Context
	)

	entities := make([]models.Entertainment, 0, 5)

	BeforeEach(func() {
		entities = dataProviderEntities()
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRepo(mockCtrl)
		cntx = context.Background()
		mockRepo.EXPECT().AddEntertainments(cntx, entities[0:2]).Return(nil).Times(1)
		mockRepo.EXPECT().AddEntertainments(cntx, entities[2:4]).Return(nil).Times(1)
		mockRepo.EXPECT().AddEntertainments(cntx, entities[4:]).Return(nil).Times(1)

		f = flusher.NewFlusher(2, mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("With valid entities slice", func() {
		It("should all save", func() {

			saver := NewSaver(cntx, 5, f)
			for _, v := range entities {
				err := saver.Save(v)

				assert.Nil(GinkgoT(), err)
			}

			saver.Close()
		})
	})

	Context("Entities greater than capacity", func() {
		It("should errors in save", func() {

			entities = append(entities, models.New(6, "title", "description"))

			saver := NewSaver(cntx, 5, f)
			for index, v := range entities {
				err := saver.Save(v)

				if index < 5 {
					assert.Nil(GinkgoT(), err)
				} else {
					assert.Error(GinkgoT(), err)
				}
			}

			saver.Close()
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
