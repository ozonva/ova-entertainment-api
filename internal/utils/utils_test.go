package utils

import (
	"reflect"
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
	if actual := Split(slice, size); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Split() = %v, want %v", expected, actual)
	}
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

	if actual, _ := Flip(list); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Flip() = %v, want %v", expected, actual)
	}
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
	if actual := Filter(slice, words); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Filter() = %v, want %v", actual, expected)
	}
}
