package models

import (
	"fmt"
	"time"

	desc "github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api/github.com/ozonva/ova-entertainment-api/pkg/ova-entertainment-api"
)

// Entertainment Основная модель микросервиса
type Entertainment struct {
	// Уникальный индитификатор сущности
	ID uint64 `db:"id"`

	// Пользователь который создаль сущность
	UserID uint64 `db:"user_id"`

	// Заголовок
	Title string

	// Подробное описание
	Description string

	// Дата события
	Date time.Time
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

// Marshall Преобразовывает модель к ответу grpc
func (e *Entertainment) Marshall() *desc.EntertainmentV1Response {
	return &desc.EntertainmentV1Response{
		ID:          e.ID,
		UserID:      e.UserID,
		Title:       e.Title,
		Description: e.Description,
		Date:        e.Date.String(),
	}
}
