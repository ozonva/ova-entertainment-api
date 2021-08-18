package models

import (
	"fmt"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	userID := uint64(2222)
	entertainment := New(userID)

	assert := testify.New(t)
	assert.Equal(entertainment.GetUserID(), userID)
}

func TestString(t *testing.T) {
	title := "Go Friday"
	description := "Description go friday"
	entertainment := New(uint64(2222))
	entertainment.SetTitle(title)
	entertainment.SetDescription(description)

	assert := testify.New(t)
	assert.Equal(entertainment.String(), fmt.Sprintf("Entertainment name: %s, description: %s", title, description))
}
