package main

import (
	"fmt"
	"strings"
)

type Table struct {
	headers                        map[int]string
	contents                       map[int]map[int]string
	maxLens                        map[int]int
	minRow, maxRow, minCol, maxCol int
}

func NewTable() *Table {
	return &Table{
		make(map[int]string),
		make(map[int]map[int]string),
		make(map[int]int),
		0, 0, 0, 0,
	}
}

func (t *Table) SetHeader(col int, header string) {
	t.headers[col] = header
	if length := len(header); length > t.maxLens[col] {
		t.maxLens[col] = length
	}
}

func (t *Table) GetHeader(col int) string {
	header, ok := t.headers[col]
	if !ok {
		return ""
	}
	return header
}

func (t *Table) SetContent(row, col int, content string) {
	t2, ok := t.contents[row]
	if !ok {
		t2 = make(map[int]string)
		t.contents[row] = t2
	}
	t2[col] = content
	if length := len(content); length > t.maxLens[col] {
		t.maxLens[col] = length
	}
	if row < t.minRow {
		t.minRow = row
	}
	if row > t.maxRow {
		t.maxRow = row
	}
	if col < t.minCol {
		t.minCol = col
	}
	if col > t.maxCol {
		t.maxCol = col
	}
}

func (t *Table) GetContent(row, col int) string {
	t2, ok := t.contents[row]
	if !ok {
		return ""
	}
	content, ok := t2[col]
	if !ok {
		return ""
	}
	return content
}

func (t *Table) GetMaxLen(col int) int {
	maxLen, ok := t.maxLens[col]
	if !ok {
		return 0
	}
	return maxLen
}

func (t *Table) Print() {
	for col, sep := t.minCol, ""; col <= t.maxCol; col++ {
		fmt.Printf("%s%*s", sep, -t.GetMaxLen(col), t.GetHeader(col))
		sep = " | "
	}
	fmt.Println()
	for col, sep := t.minCol, ""; col <= t.maxCol; col++ {
		fmt.Printf("%s%s", sep, strings.Repeat("-", t.GetMaxLen(col)))
		sep = " | "
	}
	fmt.Println()
	for row := t.minRow; row <= t.maxRow; row++ {
		for col, sep := t.minCol, ""; col <= t.maxCol; col++ {
			fmt.Printf("%s%*s", sep, -t.GetMaxLen(col), t.GetContent(row, col))
			sep = " | "
		}
		fmt.Println()
	}
	fmt.Println()
}
