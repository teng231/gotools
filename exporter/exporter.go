package exporter

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Exporter(headers []string, dataRows [][]interface{}) ([][]string, error) {
	exporter := make([][]string, 0)
	exporter = append(exporter, headers)
	for _, rows := range dataRows {
		csvRows := []string{}
		for _, row := range rows {
			if row == nil {
				csvRows = append(csvRows, "")
			}
			csvRows = append(csvRows, fmt.Sprintf("%v", row))
		}
		exporter = append(exporter, csvRows)
	}
	return exporter, nil
}

func MakeCsvFile(filename string, data [][]string) (*os.File, func()) {
	file, _ := os.Create(filename + ".csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			log.Print("Cannot write to file", err)
		}
	}
	return file, func() {
		if err := os.Remove(filename + ".csv"); err != nil {
			log.Print(err)
		}
	}
}
