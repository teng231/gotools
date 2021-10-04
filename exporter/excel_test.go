package exporter

import "testing"

func TestExcelReader(t *testing.T) {
	ExcelReader("products.xlsx", "Active")
}

func TestExcelWriter(t *testing.T) {
	headers := []string{"col1", "col2", "col3"}
	contents := [][]interface{}{
		{
			"te3", "x", 1.23,
		},
		{
			"te2", "p", 1,
		},
		{
			"te1", "x", false,
		},
	}
	err := ExcelWriter("tenguyen", headers, contents, "done.xlsx")
	if err != nil {
		t.Fail()
	}
}

func TestExcelWritermiltipleSheet(t *testing.T) {
	contents := [][]interface{}{
		{
			"te3", "x", 1.23,
		},
		{
			"te2", "p", 1,
		},
		{
			"te1", "x", false,
		},
	}
	err := ExcelWriterMultipleSheets([]*ExcelSheet{
		{Header: []string{"col1", "col2", "col3"}, Sheet: "SSS1", Rows: contents},
		{Header: []string{"col4", "col5", "col6"}, Sheet: "SSS2", Rows: contents},
	}, "done2.xlsx")
	if err != nil {
		t.Fail()
	}
}
