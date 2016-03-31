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
