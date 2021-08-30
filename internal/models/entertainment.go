package models

import (
	"fmt"
	"time"

	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
)

type Entertainment struct {
	ID          uint64 `db:"id"`
	UserID      uint64 `db:"user_id"`
	Title       string
	Description string
	Date        time.Time
}

func New(UserID uint64, title string, description string) Entertainment {
	entity := Entertainment{
		UserID:      UserID,
		Title:       title,
		Description: description,
		Date:        time.Now().Truncate(24*time.Hour).AddDate(0, 0, 7),
	}

	return entity
}

func (e *Entertainment) String() string {
	return fmt.Sprintf("Entertainment name: %s, Description: %s", e.Title, e.Description)
}

// @todo должно быть какбудто не тут
func (e *Entertainment) Marshall() *desc.EntertainmentV1Response {
	return &desc.EntertainmentV1Response{
		ID:          e.ID,
		UserID:      e.UserID,
		Title:       e.Title,
		Description: e.Description,
		Date:        e.Date.String(),
	}
}
