package main

import (
	"github.com/asukakenji/clash-royale/lib"

	"fmt"
)

const (
	fixedHeaderWidth        = 16
	upgradableHeaderWidth   = 24
	fixedContentsWidth      = 5
	upgradableContentsWidth = 7
)

var (
	fixedColHeaders      = []string{"Attribute", "Value"}
	upgradableColHeaders = []string{"Attribute", "LV1", "LV2", "LV3", "LV4", "LV5", "LV6", "LV7", "LV8", "LV9", "LV10", "LV11", "LV12"}
)

func main() {
	sep := ""
	lib.ForEachType(func(_type *lib.Type) {
		// Card Type (Troops, Buildings, Spells)
		fmt.Printf(sep)
		fmt.Printf("## %s\n", _type)
		fmt.Println()
		lib.ForEachCardOfType(_type, func(card *lib.Card) {

			// Header (Card Name)
			fmt.Printf("### %s\n", card.Name())
			fmt.Println()

			// Fixed Attribute Table
			{
				rowHeaders := []string{}
				contents := [][]string{}
				card.ForEachFixedAttribute(func(attr *lib.FixedAttribute) {
					rowHeaders = append(rowHeaders, attr.String())
					contents = append(contents, []string{card.FormattedValue(attr)})
				})
				table := NewTable(map[string]interface{}{
					"headerWidth":   fixedHeaderWidth,
					"contentsWidth": fixedContentsWidth,
					"rowHeaders":    rowHeaders,
					"colHeaders":    fixedColHeaders,
					"contents":      contents,
				})
				table.SetLastColTrim(true)
				table.Print()
			}

			// Upgradable Attribute Table
			{
				rowHeaders := []string{}
				contents := [][]string{}
				card.ForEachUpgradableAttribute(func(attr *lib.UpgradableAttribute) {
					rowHeaders = append(rowHeaders, attr.String())
					contents = append(contents, card.FormattedValues(attr))
				})
				maxLevel := card.MaxLevel()
				table := NewTable(map[string]interface{}{
					"headerWidth":   upgradableHeaderWidth,
					"contentsWidth": upgradableContentsWidth,
					"rowHeaders":    rowHeaders,
					"colHeaders":    upgradableColHeaders[0 : maxLevel+1 : maxLevel+1],
					"contents":      contents,
				})
				table.Print()
			}
		})
		sep = "\n\n\n\n"
	})
}
