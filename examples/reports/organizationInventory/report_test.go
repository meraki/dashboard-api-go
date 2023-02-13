package main

import (
	"github.com/xuri/excelize/v2"
	"reflect"
	"testing"
)

func TestNewSpreadsheet(t *testing.T) {
	tabsData := []tabData{
		{
			Name: "Tab1",
			Rows: [][]interface{}{
				{"Row 1, Cell 1", "Row 1, Cell 2", "Row 1, Cell 3"},
				{"Row 2, Cell 1", "Row 2, Cell 2", "Row 2, Cell 3"},
			},
		},
		{
			Name: "Tab2",
			Rows: [][]interface{}{
				{"Row 1, Cell 1", "Row 1, Cell 2"},
				{"Row 2, Cell 1", "Row 2, Cell 2"},
				{"Row 3, Cell 1", "Row 3, Cell 2"},
			},
		},
	}

	f := newSpreadsheet(tabsData)

	// Check that file is not nil
	if f == nil {
		t.Error("Expected a valid file, got nil")
	}

	// Check that the number of sheets in the file matches the number of tabs
	if len(f.GetSheetMap()) != len(tabsData) {
		t.Errorf("Expected %d sheets, got %d", len(tabsData), len(f.GetSheetMap()))
	}

	// Check that the data in each tab matches the input data
	for _, tabData := range tabsData {
		sheetName := tabData.Name
		rows := tabData.Rows

		rowsFromSheet, err := f.GetRows(sheetName)
		if err != nil {
			t.Errorf("Error reading rows from sheet '%s': %v", sheetName, err)
			continue
		}

		// Check that the number of rows in the sheet matches the number of rows in the tab
		if len(rowsFromSheet) != len(rows) {
			t.Errorf("Expected %d rows in sheet '%s', got %d", len(rows), sheetName, len(rowsFromSheet))
			continue
		}

		// Check that the data in each cell matches the input data
		for j, row := range rows {
			for k, cellValue := range row {
				cellIndex, err := excelize.CoordinatesToCellName(k+1, j+1)
				if err != nil {
					t.Errorf("Error getting cell index for row %d, column %d: %v", j+1, k+1, err)
					continue
				}

				cellValueFromSheet, _ := f.GetCellValue(sheetName, cellIndex)

				if !reflect.DeepEqual(cellValue, cellValueFromSheet) {
					t.Errorf("Expected cell %s in sheet '%s' to have value %v, got %v", cellIndex, sheetName, cellValue, cellValueFromSheet)
				}
			}
		}
	}
}

func TestAddRowDataToTab(t *testing.T) {
	tab := &tabData{
		Name: "TestTab",
		Rows: [][]interface{}{
			{1, "foo", true},
			{2, "bar", false},
		},
	}

	data := []interface{}{3, "baz", true}
	tab = addRowDataToTab(tab, data)

	expectedRows := [][]interface{}{
		{1, "foo", true},
		{2, "bar", false},
		{3, "baz", true},
	}

	if !reflect.DeepEqual(tab.Rows, expectedRows) {
		t.Errorf("addRowDataToTab failed. Expected: %v, actual: %v", expectedRows, tab.Rows)
	}
}
