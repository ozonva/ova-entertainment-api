package repo

import "github.com/ozonva/ova-entertainment-api/internal/models"

type Repo interface {
	AddEntities(entities []models.Entertainment) error
	ListEntities(limit, offset uint64) ([]models.Entertainment, error)
	DescribeEntity(entityId uint64) (*models.Entertainment, error)
}
