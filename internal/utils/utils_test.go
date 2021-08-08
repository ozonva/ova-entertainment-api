package utils

import (
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	expected := [][]int{
		{1, 2},
		{3, 4},
		{5},
	}
	slice := []int{1, 2, 3, 4, 5}
	size := 2

	actual := Split(slice, size)

	assert := testify.New(t)
	assert.Equal(actual, expected)
}

func TestFlip(t *testing.T) {
	list := map[int]string{
		1: "one",
		2: "two",
	}
	expected := map[string]int{
		"one": 1,
		"two": 2,
	}

	actual, _ := Flip(list)

	assert := testify.New(t)
	assert.Equal(actual, expected)
}

func TestFilter(t *testing.T) {
	slice := []string{
		"one",
		"two",
		"four",
		"five",
		"six",
		"seven",
	}
	words := []string{
		"two",
		"four",
	}
	expected := []string{
		"one",
		"five",
		"six",
		"seven",
	}

	actual := Filter(slice, words)

	assert := testify.New(t)
	assert.Equal(actual, expected)
}
