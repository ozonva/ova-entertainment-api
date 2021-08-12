package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(fileName string) func() {
	return func() {
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
			value, errors := reader.Read()
			if errors != nil {
				break
			}
			fmt.Println(value)
		}
	}
}

func main() {

	for i := 0; i < 5; i++ {
		readCsv := ReadCsvFile("file.csv")
		readCsv()
	}

}
