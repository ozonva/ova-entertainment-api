package flusher

import (
	"context"
	"errors"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	"github.com/ozonva/ova-entertainment-api/internal/utils"
)

// Flusher Сохранение моделей батчами
type Flusher interface {
	Flush(ctx context.Context, entities []models.Entertainment) error
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

func (f *flusher) Flush(ctx context.Context, entities []models.Entertainment) error {
	chunks := utils.SplitToBulks(entities, uint(f.chunkSize))

	notSavedModels := make([]models.Entertainment, 0, len(entities))
	for _, chunk := range chunks {
		err := f.entityRepo.AddEntertainments(ctx, chunk)
		if err != nil {
			return err
		}
	}

	if len(notSavedModels) == 0 {
		return nil
	}

	return errors.New("Not all saved")
}
