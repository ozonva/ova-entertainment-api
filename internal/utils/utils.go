package utils

import (
	"errors"
	"math"
)

func Split(slice []int, size int) [][]int {
	resultSliceSize := int(math.Ceil(float64(len(slice)) / float64(size)))
	result := make([][]int, resultSliceSize)

	lengthSlice := len(slice)
	start := 0
	end := size

	for index := 0; index < resultSliceSize; index++ {

		if end > lengthSlice {
			end = lengthSlice
		}

		result[index] = append(result[index], slice[start:end]...)

		start += size
		end += size
	}

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

	index := 0
	var exist bool
	for _, value := range slice {

		exist = false
		for _, word := range words {
			if word == value {
				exist = true
				break
			}
		}

		if !exist {
			slice[index] = value
			index++
		}
	}

	return slice[:index]
}
