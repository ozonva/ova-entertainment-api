package models

import (
	"fmt"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	userID := uint64(2222)
	entertainment := New(userID, "title", "description")

	assert := testify.New(t)
	assert.Equal(entertainment.UserID, userID)
}

func TestString(t *testing.T) {
	title := "Go Friday"
	description := "Description go friday"
	entertainment := New(uint64(2222), "title", "description")
	entertainment.Title = title
	entertainment.Description = description

	assert := testify.New(t)
	assert.Equal(entertainment.String(), fmt.Sprintf("Entertainment name: %s, Description: %s", title, description))
}
