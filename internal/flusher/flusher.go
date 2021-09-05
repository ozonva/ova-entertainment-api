package flusher

import (
	"errors"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	"github.com/ozonva/ova-entertainment-api/internal/utils"
)

// Flusher Сохранение моделей батчами
type Flusher interface {
	Flush(entities []models.Entertainment) error
}

func NewFlusher(chunkSize int, entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (f *flusher) Flush(entities []models.Entertainment) error {
	chunks := utils.SplitToBulks(entities, uint(f.chunkSize))

	notSavedModels := make([]models.Entertainment, 0, len(entities))
	for _, chunk := range chunks {
		err := f.entityRepo.AddEntertainments(chunk)
		if err != nil {
			return err
		}
	}

	if len(notSavedModels) == 0 {
		return nil
	}

	return errors.New("Not all saved")
}
