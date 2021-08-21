package saver

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-entertainment-api/internal/flusher"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	"github.com/stretchr/testify/assert"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *repo.MockRepo
		f        flusher.Flusher
	)

	entities := make([]models.Entertainment, 0, 5)

	BeforeEach(func() {
		entities = dataProviderEntities()
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRepo(mockCtrl)
		mockRepo.EXPECT().AddEntities(entities[0:2]).Return(nil).Times(1)
		mockRepo.EXPECT().AddEntities(entities[2:4]).Return(nil).Times(1)
		mockRepo.EXPECT().AddEntities(entities[4:]).Return(nil).Times(1)

		f = flusher.NewFlusher(2, mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("With valid entities slice", func() {
		It("should all save", func() {

			saver := NewSaver(5, f)
			saver.Init()
			for _, v := range entities {
				err := saver.Save(v)

				assert.Nil(GinkgoT(), err)
			}

			time.Sleep(5 * time.Second)
			saver.Close()
			time.Sleep(2 * time.Second)
		})
	})

	Context("Entities greater than capacity", func() {
		It("should errors in save", func() {

			entities = append(entities, models.New(6))

			saver := NewSaver(5, f)
			saver.Init()
			for index, v := range entities {
				err := saver.Save(v)

				if index < 5 {
					assert.Nil(GinkgoT(), err)
				} else {
					assert.Error(GinkgoT(), err)
				}
			}

			time.Sleep(5 * time.Second)
			saver.Close()
			time.Sleep(2 * time.Second)
		})
	})
})

func dataProviderEntities() []models.Entertainment {
	entertainment1 := models.New(1)
	entertainment2 := models.New(2)
	entertainment3 := models.New(3)
	entertainment4 := models.New(4)
	entertainment5 := models.New(5)

	slice := []models.Entertainment{
		entertainment1,
		entertainment2,
		entertainment3,
		entertainment4,
		entertainment5,
	}

	return slice
}
