package main

import (
	"github.com/asukakenji/clash-royale/attribute"
	"github.com/asukakenji/clash-royale/card"
	"github.com/asukakenji/clash-royale/Type"

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
	Type.ForEach(func(t Type.Type) {
		// Card Type (Troops, Buildings, Spells)
		fmt.Printf(sep)
		fmt.Printf("## %s\n", t)
		fmt.Println()
		card.ForEachCardOfType(t, func(c *card.Card) {

			// Header (Card Name)
			fmt.Printf("### %s\n", c.Name())
			fmt.Println()

			// Fixed Attribute Table
			{
				rowHeaders := []string{}
				contents := [][]string{}
				c.ForEachFixedAttribute(func(a attribute.Fixed) {
					rowHeaders = append(rowHeaders, a.String())
					contents = append(contents, []string{c.FormattedValue(a)})
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
				c.ForEachUpgradableAttribute(func(a attribute.Upgradable) {
					rowHeaders = append(rowHeaders, a.String())
					contents = append(contents, c.FormattedValues(a))
				})
				maxLevel := c.MaxLevel()
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
