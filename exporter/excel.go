package exporter

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/xuri/excelize/v2"
)

func ExcelReader(path, sheet string) ([]string, [][]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, nil, err
	}
	if f == nil {
		return nil, nil, errors.New("not found file excel")
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, nil, errors.New("not found excel rows")
	}
	headers := rows[0]
	contents := make([][]string, 0)
	if len(rows) > 1 {
		contents = rows[1:]
	}
	return headers, contents, nil
}
func ExcelReaderStream(file io.Reader, sheet string) ([]string, [][]string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, nil, err
	}
	if f == nil {
		return nil, nil, errors.New("not found file excel")
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, nil, errors.New("not found excel rows")
	}
	// headers := make([]string, 0)
	headers := rows[0]
	contents := make([][]string, 0)
	if len(rows) > 1 {
		contents = rows[1:]
	}
	return headers, contents, nil
}

var (
	axisX = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func setCellByFormat(f *excelize.File, sheet, axis string, x interface{}) {
	switch x.(type) {
	case nil:
		f.SetCellValue(sheet, axis, "")
	case float64:
		f.SetCellFloat(sheet, axis, x.(float64), 10, 64)
	case float32:
		f.SetCellFloat(sheet, axis, float64(x.(float32)), 10, 64)
	case int:
		f.SetCellInt(sheet, axis, x.(int))
	case int32:
		f.SetCellInt(sheet, axis, int(x.(int32)))
	case int64:
		f.SetCellInt(sheet, axis, int(x.(int64)))
	case string:
		f.SetCellStr(sheet, axis, x.(string))
	case bool:
		f.SetCellBool(sheet, axis, x.(bool))
	default:
		f.SetCellValue(sheet, axis, x)
	}
}
func ExcelWriter(sheet string, headers []string, rows [][]interface{}, fileout string) error {
	f := excelize.NewFile()
	// Create a new worksheet.
	idx := f.NewSheet(sheet)
	if len(headers) > 0 {
		// default i = 1
		for i, axis := range axisX[0:len(headers)] {
			axisPoint := fmt.Sprintf("%s%d", axis, 1)
			f.SetCellValue(sheet, axisPoint, headers[i])
			f.SetCellRichText(sheet, axisPoint, []excelize.RichTextRun{
				{
					Text: headers[i],
					Font: &excelize.Font{
						Bold: true,
					},
				},
			})
		}
	}
	axisXContent := axisX[0:len(rows[0])]
	log.Print(axisXContent)
	for i, row := range rows {
		for j, ax := range axisXContent {
			axisPoint := fmt.Sprintf("%s%d", ax, i+2)
			setCellByFormat(f, sheet, axisPoint, row[j])
		}
	}
	f.SetActiveSheet(idx)
	return f.SaveAs(fileout)
}

type ExcelSheet struct {
	Sheet  string
	Header []string
	Rows   [][]interface{}
}

func ExcelWriterMultipleSheets(excelSheets []*ExcelSheet, fileout string) error {
	f := excelize.NewFile()
	for _, excelSheet := range excelSheets {
		// Create a new worksheet.
		f.NewSheet(excelSheet.Sheet)
		if len(excelSheet.Header) > 0 {
			// default i = 1
			for i, axis := range axisX[0:len(excelSheet.Header)] {
				axisPoint := fmt.Sprintf("%s%d", axis, 1)
				f.SetCellValue(excelSheet.Sheet, axisPoint, excelSheet.Header[i])
				f.SetCellRichText(excelSheet.Sheet, axisPoint, []excelize.RichTextRun{
					{
						Text: excelSheet.Header[i],
						Font: &excelize.Font{
							Bold: true,
						},
					},
				})
			}
		}
		if len(excelSheet.Rows) == 0 {
			continue
		}
		axisXContent := axisX[0:len(excelSheet.Rows[0])]
		log.Print(axisXContent)
		for i, row := range excelSheet.Rows {
			for j, ax := range axisXContent {
				axisPoint := fmt.Sprintf("%s%d", ax, i+2)
				setCellByFormat(f, excelSheet.Sheet, axisPoint, row[j])
			}
		}
	}
	return f.SaveAs(fileout)
}
