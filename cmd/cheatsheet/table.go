package main

import (
	"fmt"
	"strings"
)

type Alignment int

const (
	LEFT  Alignment = -1
	RIGHT Alignment = 1
)

type Table struct {
	rowHeaders         []string
	colHeaders         []string
	contents           [][]string
	headerWidth        int
	contentsWidth      int
	rowHeaderAlignment Alignment
	colHeaderAlignment Alignment
	contentAlignment   Alignment
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
		LEFT,
		LEFT,
		RIGHT,
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

func (t *Table) SetContents(contents [][]string) {
	t.contents = contents
}

func (t *Table) SetRowContents(row int, rowContents []string) {
	t.contents[row] = rowContents
}

func (t *Table) SetCell(row, col int, value string) {
	for row >= len(t.contents) {
		t.contents = append(t.contents, []string{})
	}
	for col >= len(t.contents[row]) {
		t.contents[row] = append(t.contents[row], "")
	}
	t.contents[row][col] = value
}

func (t *Table) SetRowHeaderAlignment(a Alignment) {
	t.rowHeaderAlignment = a
}

func (t *Table) SetColHeaderAlignment(a Alignment) {
	t.colHeaderAlignment = a
}

func (t *Table) SetContentAlignment(a Alignment) {
	t.contentAlignment = a
}

func (t *Table) Print() {
	limit := len(t.colHeaders) - 2

	// Table header
	fmt.Printf("%*s", int(t.rowHeaderAlignment)*t.headerWidth, t.colHeaders[0])
	for i, colHeader := range t.colHeaders[1:] {
		if i != limit {
			fmt.Printf(" | %*s", int(t.colHeaderAlignment)*t.contentsWidth, colHeader)
		} else {
			s := fmt.Sprintf(" | %*s", int(t.colHeaderAlignment)*t.contentsWidth, colHeader)
			fmt.Printf(strings.TrimRight(s, " "))
		}
	}
	fmt.Println()

	// Header separator
	fmt.Printf("%*s", int(t.rowHeaderAlignment)*t.headerWidth, strings.Repeat("-", t.headerWidth))
	colSeparator := strings.Repeat("-", t.contentsWidth)
	for i := range t.colHeaders[1:] {
		if i != limit {
			fmt.Printf(" | %*s", int(t.colHeaderAlignment)*t.contentsWidth, colSeparator)
		} else {
			s := fmt.Sprintf(" | %*s", int(t.colHeaderAlignment)*t.contentsWidth, colSeparator)
			fmt.Printf(strings.TrimRight(s, " "))
		}
	}
	fmt.Println()

	// Table contents
	for row, rowHeader := range t.rowHeaders {
		fmt.Printf("%*s", int(t.rowHeaderAlignment)*t.headerWidth, rowHeader)
		for i, content := range t.contents[row] {
			if i != limit {
				fmt.Printf(" | %*s", int(t.contentAlignment)*t.contentsWidth, content)
			} else {
				s := fmt.Sprintf(" | %*s", int(t.contentAlignment)*t.contentsWidth, content)
				fmt.Printf(strings.TrimRight(s, " "))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
