package entities

import (
	"fmt"
	"time"
)

type Entertainment struct {
	UserID      uint64
	Title       string
	Description string
	Date        time.Time
}

func New(UserID uint64) Entertainment {
	entity := Entertainment{
		UserID: UserID,
		Date:   time.Now().Truncate(24*time.Hour).AddDate(0, 0, 7),
	}

	return entity
}

func (e *Entertainment) String() string {
	return fmt.Sprintf("Entertainment name: %s, description: %s", e.Title, e.Description)
}

func (e *Entertainment) SetTitle(title string) {
	e.Title = title
}

func (e Entertainment) GetTitle() string {
	return e.Title
}

func (e *Entertainment) SetDate(date time.Time) {
	e.Date = date
}

func (e Entertainment) GetDate() time.Time {
	return e.Date
}

func (e *Entertainment) SetDescription(description string) {
	e.Description = description
}

func (e Entertainment) GetDescription() string {
	return e.Description
}
