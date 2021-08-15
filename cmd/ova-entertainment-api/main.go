package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	reader := csv.NewReader(file)
	for {
		value, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(value)
	}
}

func main() {

	for i := 0; i < 5; i++ {
		ReadCsvFile("file.csv")
	}

}
