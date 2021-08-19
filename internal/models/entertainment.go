package models

import (
	"fmt"
	"time"
)

type Entertainment struct {
	userID      uint64
	title       string
	description string
	date        time.Time
}

func New(UserID uint64) Entertainment {
	entity := Entertainment{
		userID: UserID,
		date:   time.Now().Truncate(24*time.Hour).AddDate(0, 0, 7),
	}

	return entity
}

func (e Entertainment) GetUserID() uint64 {
	return e.userID
}

func (e *Entertainment) String() string {
	return fmt.Sprintf("Entertainment name: %s, description: %s", e.title, e.description)
}

func (e *Entertainment) SetTitle(title string) {
	e.title = title
}

func (e Entertainment) GetTitle() string {
	return e.title
}

func (e *Entertainment) SetDate(date time.Time) {
	e.date = date
}

func (e Entertainment) GetDate() time.Time {
	return e.date
}

func (e *Entertainment) SetDescription(description string) {
	e.description = description
}

func (e Entertainment) GetDescription() string {
	return e.description
}
