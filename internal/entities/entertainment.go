package entities

import (
	"fmt"
	"time"
)

type Entertainment struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Date        time.Time
}

func (e *Entertainment) New(UserID int64) *Entertainment {
	entity := &Entertainment{
		UserID: UserID,
		Date:   time.Now().AddDate(0, 0, 7),
	}

	return entity
}

func (e *Entertainment) String() {
	fmt.Printf("Entertainment name: %s, description: %s", e.Title, e.Description)
}

func (e *Entertainment) SetTitle(title string) {
	e.Title = title
}

func (e *Entertainment) SetDate(date time.Time) {
	e.Date = date
}

func (e *Entertainment) SetDescription(description string) {
	e.Description = description
}
