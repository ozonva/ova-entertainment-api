package utils

func Split(slice []int, size int) [][]int {
	result := [][]int{}
	length := len(slice)
	for index := 0; index < length; index += size {

		end := index + size
		if end > length {
			end = length
		}

		result = append(result, slice[index:end])
	}

	return result
}

func Flip(list map[int]string) map[string]int {
	result := make(map[string]int)

	for key, value := range list {
		result[value] = key
	}

	return result
}

func Filter(slice []string, words []string) []string {

	result := []string{}
	isExistInWords := false
	for _, valueSlice := range slice {
		isExistInWords = false
		for _, valueWord := range words {
			if valueSlice == valueWord {
				isExistInWords = true
			}
		}

		if !isExistInWords {
			result = append(result, valueSlice)
		}

	}

	return result
}
