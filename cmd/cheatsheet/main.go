package main

import (
	"github.com/asukakenji/clash-royale/lib"

	"fmt"
	"strings"
)

const (
	attrTitle     = "Attribute"
	valueTitle    = "Value"
	attrTitleLen  = len(attrTitle)
	valueTitleLen = len(valueTitle)
	attrValueLen  = 7
)

var (
	fixedColHeaders      = []string{"Attribute", "Value"}
	upgradableColHeaders = []string{"Attribute", "LV1", "LV2", "LV3", "LV4", "LV5", "LV6", "LV7", "LV8", "LV9", "LV10", "LV11", "LV12"}
)

func main() {
	fixedHeaderWidth := 16
	upgradableHeaderWidth := 24
	fixedContentsWidth := 5
	upgradableContentsWidth := 7

	emptyType := lib.Type("")
	_type := emptyType

	for _, card := range lib.CARDS {
		// Card Type (Troops, Buildings, Spells)
		if cardType := card.GetValue(lib.TYPE).(lib.Type); cardType != _type {
			if _type != emptyType {
				fmt.Printf("\n\n\n\n")
			}
			fmt.Printf("## %s\n", cardType)
			fmt.Println()
			_type = cardType
		}

		// Header (Card Name)
		fmt.Printf("### %s\n", card.GetValue(lib.NAME))
		fmt.Println()

		// Fixed Attribute Table
		{
			fas := card.GetFixedAttributes()
			rh := make([]string, len(fas))
			for row, fa := range fas {
				rh[row] = fa.String()
			}
			fat := NewTable(map[string]interface{}{
				"headerWidth":   fixedHeaderWidth,
				"contentsWidth": fixedContentsWidth,
				"rowHeaders":    rh,
				"colHeaders":    fixedColHeaders,
			})
			fat.SetLastColTrim(true)
			for row, fa := range fas {
				fv := card.GetFormattedValue(fa)
				fat.SetRowContents(row, []string{fv})
			}
			fat.Print()
		}

		// Upgradable Attribute Table
		{
			uas := card.GetUpgradableAttributes()
			rh := make([]string, len(uas))
			for row, ua := range uas {
				rh[row] = ua.String()
			}
			maxLevel := card.GetMaxLevel()
			uat := NewTable(map[string]interface{}{
				"headerWidth":   upgradableHeaderWidth,
				"contentsWidth": upgradableContentsWidth,
				"rowHeaders":    rh,
				"colHeaders":    upgradableColHeaders[0 : maxLevel+1 : maxLevel+1],
			})
			for row, ua := range uas {
				fvs := card.GetFormattedValues(ua)
				uat.SetRowContents(row, fvs)
			}
			uat.Print()
		}
	}
}

func main1() {
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
