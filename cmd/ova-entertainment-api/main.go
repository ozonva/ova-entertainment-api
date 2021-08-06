package main

import (
	"fmt"
	"reflect"

	"github.com/ozonva/ova-entertainment-api/internal/utils"
)

func main() {

	expected := [][]int{
		{1, 2},
		{3, 4},
		{5},
	}
	slice := []int{1, 2, 3, 4, 5}
	size := 2
	if actual := utils.Split(slice, size); !reflect.DeepEqual(expected, actual) {
		fmt.Println("Debug")
	}

	fmt.Println("Hello from ova-entertainment-api")
}
