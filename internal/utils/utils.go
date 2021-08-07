package utils

import (
	"errors"
)

func Split(slice []int, size int) [][]int {

	result := make([][]int, 0, (len(slice)+size-1)/size)

	for size < len(slice) {
		result = append(result, slice[0:size])
		slice = slice[size:]
	}
	result = append(result, slice)

	return result
}

func Flip(list map[int]string) (map[string]int, error) {
	result := make(map[string]int)

	for key, value := range list {
		if _, ok := result[value]; ok {
			return nil, errors.New("Not unique value")
		}
		result[value] = key
	}

	return result, nil
}

func Filter(slice []string, words []string) []string {

	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}

	result := make([]string, len(slice))
	copy(result, slice)

	index := 0
	for _, value := range slice {
		if _, found := wordMap[value]; !found {
			result[index] = value
			index++
			continue
		}
	}

	return result[:index]
}
