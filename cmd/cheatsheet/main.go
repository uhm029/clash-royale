package main

import (
	"github.com/asukakenji/clash-royale/lib"

	"fmt"
	"strings"
)

// --- Table ---

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

/*
func main() {
	t := NewTable()
	t.SetHeader(0, "Attribute")
	t.SetHeader(1, "Value")
	t.SetContent(0, 0, "Name")
	t.SetContent(0, 1, "Knight")
	t.Print()
}
*/

// --- Main ---

const (
	attrTitle     = "Attribute"
	valueTitle    = "Value"
	attrTitleLen  = len(attrTitle)
	valueTitleLen = len(valueTitle)
	attrValueLen  = 7
)

func main() {
	fixedAttrNameLen := attrTitleLen
	upgradableAttrNameLen := attrTitleLen
	for _, attr := range lib.ATTRIBUTES {
		attrNameLen := len(attr.String())
		switch attr.(type) {
		case *lib.FixedAttribute:
			if attrNameLen > fixedAttrNameLen {
				fixedAttrNameLen = attrNameLen
			}
		case *lib.UpgradableAttribute:
			if attrNameLen > upgradableAttrNameLen {
				upgradableAttrNameLen = attrNameLen
			}
		}
	}

	var emptyType = lib.Type("")
	var _type lib.Type = emptyType

	for _, card := range lib.CARDS {
		if cardType := card.GetValue(lib.TYPE).(lib.Type); cardType != _type {
			if _type != emptyType {
				fmt.Printf("\n\n\n\n")
			}
			fmt.Printf("## %s\n", cardType)
			fmt.Println()
			_type = cardType
		}

		// Header
		fmt.Printf("### %s\n", card.GetValue(lib.NAME))
		fmt.Println()

		// Fixed Attribute Table
		fmt.Printf("%*s | %s\n", -fixedAttrNameLen, attrTitle, valueTitle)
		fmt.Printf("%*s | %s\n", fixedAttrNameLen, strings.Repeat("-", fixedAttrNameLen), strings.Repeat("-", valueTitleLen))
		for _, attr := range lib.ATTRIBUTES {
			fattr, ok := attr.(*lib.FixedAttribute)
			if !ok {
				continue
			}
			if value := card.GetFormattedValue(fattr); value != "" {
				fmt.Printf("%*s | %s\n", -fixedAttrNameLen, fattr, value)
			}
		}
		fmt.Println()

		// Upgradable Attribute Table
		// Header 1
		fmt.Printf("%*s", -upgradableAttrNameLen, attrTitle)
		// Any field will do, not just "cards".
		for level := range card.GetFormattedValues(lib.CARDS_REQ) {
			fmt.Printf(" | %*s", -attrValueLen, fmt.Sprintf("LV%d", level+1))
		}
		fmt.Println()

		// Header 2
		fmt.Printf("%*s", upgradableAttrNameLen, strings.Repeat("-", upgradableAttrNameLen))
		// Any field will do, not just "cards".
		for range card.GetFormattedValues(lib.CARDS_REQ) {
			fmt.Printf(" | %*s", attrValueLen, strings.Repeat("-", attrValueLen))
		}
		fmt.Println()

		// Content
		for _, attr := range lib.ATTRIBUTES {
			uattr, ok := attr.(*lib.UpgradableAttribute)
			if !ok {
				continue
			}
			if values := card.GetFormattedValues(uattr); values != nil {
				fmt.Printf("%*s", -upgradableAttrNameLen, uattr)
				for _, value := range values {
					fmt.Printf(" | %*s", attrValueLen, value)
				}
				fmt.Println()
			}
		}
		fmt.Println()
	}
}
