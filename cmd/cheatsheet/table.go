package main

import (
	"fmt"
	"strings"
)

type Table struct {
	rowHeaders    []string
	colHeaders    []string
	contents      [][]string
	headerWidth   int
	contentsWidth int
	isLastColTrim bool
}

func NewTable(options map[string]interface{}) *Table {
	// Column Count
	//cc, ok := options["colCount"]
	//if !ok {
	//	panic()
	//}
	//colCount, ok := cc.(int)
	//if !ok {
	//	panic()
	//}
	// Header Width
	headerWidth := 0
	hw, ok := options["headerWidth"]
	if ok {
		headerWidth = hw.(int)
	}
	// Contents Width
	contentsWidth := 0
	cw, ok := options["contentsWidth"]
	if ok {
		contentsWidth = cw.(int)
	}
	// Row Headers
	var rowHeaders []string = nil
	rh, ok := options["rowHeaders"]
	if ok {
		rowHeaders = rh.([]string)
	}
	// Column Headers
	var colHeaders []string = nil
	ch, ok := options["colHeaders"]
	if ok {
		colHeaders = ch.([]string)
	}
	// Contents
	var contents [][]string = nil
	c, ok := options["contents"]
	if ok {
		contents = c.([][]string)
	} else {
		// Row Count
		rowCount := 0
		if rowHeaders != nil {
			rowCount = len(rowHeaders)
		}
		rc, ok := options["rowCount"]
		if ok {
			rowCount = rc.(int)
		}
		contents = make([][]string, rowCount)
	}
	//for row := range contents {
	//	contents[row] = make([]string, colCount)
	//}
	return &Table{
		rowHeaders,
		colHeaders,
		contents,
		headerWidth,
		contentsWidth,
		false,
	}
}

func (t *Table) SetHeaderWidth(width int) {
	t.headerWidth = width
}

func (t *Table) SetContentsWidth(width int) {
	t.contentsWidth = width
}

func (t *Table) SetRowHeaders(rowHeaders []string) {
	t.rowHeaders = rowHeaders
}

func (t *Table) SetRowHeader(row int, rowHeader string) {
	t.rowHeaders[row] = rowHeader
}

func (t *Table) SetColHeaders(colHeaders []string) {
	t.colHeaders = colHeaders
}

func (t *Table) SetRowContents(row int, rowContents []string) {
	t.contents[row] = rowContents
}

func (t *Table) SetLastColTrim(isLastColTrim bool) {
	t.isLastColTrim = isLastColTrim
}

func (t *Table) Print() {
	sep := " | "

	// Table header
	fmt.Printf("%*s", -t.headerWidth, t.colHeaders[0])
	for _, colHeader := range t.colHeaders[1:] {
		fmt.Printf("%s%*s", sep, -t.contentsWidth, colHeader)
	}
	fmt.Println()

	// Header separator
	fmt.Printf("%*s", -t.headerWidth, strings.Repeat("-", t.headerWidth))
	for range t.colHeaders[1:] {
		fmt.Printf("%s%*s", sep, -t.contentsWidth, strings.Repeat("-", t.contentsWidth))
	}
	fmt.Println()

	// Table contents
	if t.isLastColTrim {
		for row, rowHeader := range t.rowHeaders {
			fmt.Printf("%*s", -t.headerWidth, rowHeader)
			col := 0
			colX := len(t.contents[row]) - 1
			for col = 0; col < colX; col++ {
				fmt.Printf("%s%*s", sep, t.contentsWidth, t.contents[row][col])
			}
			fmt.Printf("%s%s\n", sep, t.contents[row][col])
		}
	} else {
		for row, rowHeader := range t.rowHeaders {
			fmt.Printf("%*s", -t.headerWidth, rowHeader)
			for _, content := range t.contents[row] {
				fmt.Printf("%s%*s", sep, t.contentsWidth, content)
			}
			fmt.Println()
		}
	}
	fmt.Println()
}
