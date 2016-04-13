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
	var _type *lib.Type = nil

	for _, card := range lib.CARDS {
		// Card Type (Troops, Buildings, Spells)
		if cardType := card.GetType(); cardType != _type {
			if _type != nil {
				fmt.Printf("\n\n\n\n")
			}
			fmt.Printf("## %s\n", cardType)
			fmt.Println()
			_type = cardType
		}

		// Header (Card Name)
		fmt.Printf("### %s\n", card.GetName())
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
