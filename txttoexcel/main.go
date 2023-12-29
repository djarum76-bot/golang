package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	txt := [2]string{"data-running-1676049202608", "data-running-1676047979830"}

	for i := 0; i < len(txt); i++ {
		// Open the text file
		txtFile := fmt.Sprintf("%v.txt", txt[i])
		xlsxFile := fmt.Sprintf("%v.xlsx", txt[i])

		file, err := os.Open(txtFile)
		if err != nil {
			log.Fatalf("failed to open text file: %s", err)
		}
		defer file.Close()

		// Create a new Excel file
		excelFile := xlsx.NewFile()
		sheet, err := excelFile.AddSheet("Sheet1")
		if err != nil {
			log.Fatalf("failed to add sheet: %s", err)
		}

		// Read the text file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, ",")
			row := sheet.AddRow()
			for _, field := range fields {
				cell := row.AddCell()
				cell.Value = field
			}
		}

		// Save the Excel file
		err = excelFile.Save(xlsxFile)
		if err != nil {
			log.Fatalf("failed to save Excel file: %s", err)
		}

		fmt.Println("Text file successfully converted to Excel file!")
	}
}
