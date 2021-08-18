package utils

import (
	"errors"
	"github.com/ozonva/ova-entertainment-api/internal/models"
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

	wordMap := make(map[string]struct{})
	for _, word := range words {
		wordMap[word] = struct{}{}
	}

	result := make([]string, len(slice))
	copy(result, slice)

	index := 0
	for _, value := range slice {
		if _, found := wordMap[value]; !found {
			result[index] = value
			index++
		}
	}

	return result[:index]
}

func SplitToBulks(slice []models.Entertainment, size uint) [][]models.Entertainment {
	batchSize := int(size)
	result := make([][]models.Entertainment, 0, (len(slice)+batchSize-1)/batchSize)
	for batchSize < len(slice) {
		result = append(result, slice[0:batchSize])
		slice = slice[batchSize:]
	}
	result = append(result, slice)

	return result
}

func SliceToMap(slice []models.Entertainment) (map[uint64]models.Entertainment, error) {
	result := make(map[uint64]models.Entertainment)

	for _, entity := range slice {
		userID := entity.GetUserID()
		if _, ok := result[userID]; ok {
			return nil, errors.New("Not unique value")
		}

		result[userID] = entity
	}

	return result, nil
}
