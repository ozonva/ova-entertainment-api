package flusher

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *repo.MockRepo
	)

	entities := dataProviderEntities()

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRepo(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("With valid entities slice", func() {
		It("should all save", func() {
			mockRepo.EXPECT().AddEntertainments(entities[0:2]).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntertainments(entities[2:4]).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntertainments(entities[4:]).Return(nil).Times(1)

			flusher := NewFlusher(2, mockRepo)
			notSaved := flusher.Flush(entities)

			assert.Nil(GinkgoT(), notSaved)
		})
	})

	Context("With some errors entities slice", func() {
		It("should return errors", func() {
			mockRepo.EXPECT().AddEntertainments(entities[0:2]).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntertainments(entities[2:4]).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntertainments(entities[4:]).Return(errors.New("Can`t save slice")).Times(1)

			flusher := NewFlusher(2, mockRepo)
			err := flusher.Flush(entities)

			assert.Error(GinkgoT(), err)
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
