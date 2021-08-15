package utils

import (
	"github.com/ozonva/ova-entertainment-api/internal/entities"
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

func TestSplitToBulks(t *testing.T) {
	entertainment1 := entities.New(1)
	entertainment2 := entities.New(2)
	entertainment3 := entities.New(3)
	entertainment4 := entities.New(4)
	entertainment5 := entities.New(5)

	slice := []entities.Entertainment{
		entertainment1,
		entertainment2,
		entertainment3,
		entertainment4,
		entertainment5,
	}

	expected := [][]entities.Entertainment{
		{
			entertainment1,
			entertainment2,
		}, {
			entertainment3,
			entertainment4,
		}, {
			entertainment5,
		},
	}
	actual := SplitToBulks(slice, uint(2))

	assert := testify.New(t)
	assert.Equal(actual, expected)
}

func TestSliceToMap(t *testing.T) {
	entertainment1 := entities.New(1)
	entertainment2 := entities.New(2)
	entertainment3 := entities.New(3)
	entertainment4 := entities.New(4)
	entertainment5 := entities.New(5)

	slice := []entities.Entertainment{
		entertainment1,
		entertainment2,
		entertainment3,
		entertainment4,
		entertainment5,
	}
	expected := map[uint64]entities.Entertainment{
		1: entertainment1,
		2: entertainment2,
		3: entertainment3,
		4: entertainment4,
		5: entertainment5,
	}

	assert := testify.New(t)

	actual, err := SliceToMap(slice)
	if err != nil {
		assert.Fail("Error != nil")
	}

	assert.Equal(actual, expected)
}
