package flusher

import (
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/ozonva/ova-entertainment-api/internal/repo"
	"github.com/ozonva/ova-entertainment-api/internal/utils"
)

type Flusher interface {
	Flush(entities []models.Entertainment) []models.Entertainment
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

func (f *flusher) Flush(entities []models.Entertainment) []models.Entertainment {
	chunks := utils.SplitToBulks(entities, uint(f.chunkSize))

	notSavedModels := make([]models.Entertainment, 0, len(entities))
	for _, chunk := range chunks {
		err := f.entityRepo.AddEntities(chunk)
		if err != nil {
			notSavedModels = append(notSavedModels, chunk...)
		}
	}

	if len(notSavedModels) == 0 {
		return nil
	}

	return notSavedModels
}
