package main

import (
	"github.com/asukakenji/clash-royale/attr"
	"github.com/asukakenji/clash-royale/card"
	"github.com/asukakenji/clash-royale/format"
	"github.com/asukakenji/clash-royale/king"
	"github.com/asukakenji/clash-royale/typ"

	"bytes"
	"fmt"
)

const (
	fixedHeaderWidth        = 16
	upgradableHeaderWidth   = 24
	fixedContentsWidth      = 5
	upgradableContentsWidth = 7
)

const (
	kKingLevelPlayerLevel  = "King Level (Player Level)"
	kRarity                = "Rarity"
	kCardsByTypeAndRarity  = "Cards (by Type and Rarity)"
	kCardsByTypeAndArena   = "Cards (by Type and Arena)"
	kCardsByRarityAndArena = "Cards (by Rarity and Arena)"
)

var (
	towerFixedRowHeaders      = []string{"Range", "Hit Speed"}
	towerUpgradableRowHeaders = []string{"Hitpoints", "Damage"}
	fixedColHeaders           = []string{"Attribute", "Value"}
	upgradableColHeaders      = []string{"Attribute", "LV1", "LV2", "LV3", "LV4", "LV5", "LV6", "LV7", "LV8", "LV9", "LV10", "LV11", "LV12", "LV13"}
)

func ilink(s string) string {
	var result bytes.Buffer

	result.WriteRune('[')
	result.WriteString(s)
	result.WriteRune(']')
	result.WriteRune('(')
	result.WriteRune('#')

	for _, runeValue := range s {
		switch {
		case 'A' <= runeValue && runeValue <= 'Z':
			// Turn uppercase letters into lowercase
			result.WriteRune(runeValue + 0x20)
		case 'a' <= runeValue && runeValue <= 'z':
			// Copy lowercase letters
			result.WriteRune(runeValue)
		default:
			switch runeValue {
			case ' ':
				// Turn spaces into dashes
				result.WriteRune('-')
			case '.', '(', ')':
				// Skip punctuations
			default:
				panic("Unknown character found: " + s)
			}
		}
	}

	result.WriteRune(')')
	return result.String()
}

func main() {
	fmt.Println("# Clash Royale Cheatsheet")
	fmt.Println()
	fmt.Println("## Index")
	fmt.Printf("- %s\n", ilink(kKingLevelPlayerLevel))
	fmt.Printf("- %s\n", ilink(kRarity))
	fmt.Printf("- %s\n", ilink(kCardsByTypeAndRarity))
	fmt.Printf("- %s\n", ilink(kCardsByTypeAndArena))
	fmt.Printf("- %s\n", ilink(kCardsByRarityAndArena))
	fmt.Printf("\n\n\n")

	{
		fmt.Printf("## %s\n", kKingLevelPlayerLevel)
		fmt.Println()
		table := NewTable(map[string]interface{}{
			"headerWidth":   19,
			"contentsWidth": 6,
			"rowHeaders":    []string{"Experience Required"},
			"colHeaders":    upgradableColHeaders,
		})
		king.ForEachLevel(func(k king.King) {
			table.SetCell(0, k.Level()-1, format.Int(k.ExperienceRequired()))
		})
		table.Print()
	}

	{
		fmt.Println("### King's Tower")
		fmt.Println()
		tableF := NewTable(map[string]interface{}{
			"headerWidth":   9,
			"contentsWidth": 6,
			"rowHeaders":    towerFixedRowHeaders,
			"colHeaders":    fixedColHeaders,
		})
		tableF.SetCell(0, 0, format.Range(king.KingTowerAtLevel(0).Range()))
		tableF.SetCell(1, 0, format.Time(king.KingTowerAtLevel(0).HitSpeed()))
		tableF.Print()
		tableU := NewTable(map[string]interface{}{
			"headerWidth":   9,
			"contentsWidth": 6,
			"rowHeaders":    towerUpgradableRowHeaders,
			"colHeaders":    upgradableColHeaders,
		})
		king.ForEachLevel(func(k king.King) {
			lv := k.Level()
			tableU.SetCell(0, lv-1, format.Int(king.KingTowerAtLevel(lv).Hitpoints()))
			tableU.SetCell(1, lv-1, format.Int(king.KingTowerAtLevel(lv).Damage()))
		})
		tableU.Print()
	}

	{
		fmt.Println("### Arena Towers")
		fmt.Println()
		tableF := NewTable(map[string]interface{}{
			"headerWidth":   9,
			"contentsWidth": 6,
			"rowHeaders":    towerFixedRowHeaders,
			"colHeaders":    fixedColHeaders,
		})
		tableF.SetCell(0, 0, format.Range(king.ArenaTowerAtLevel(0).Range()))
		tableF.SetCell(1, 0, format.Time(king.ArenaTowerAtLevel(0).HitSpeed()))
		tableF.Print()
		tableU := NewTable(map[string]interface{}{
			"headerWidth":   9,
			"contentsWidth": 6,
			"rowHeaders":    towerUpgradableRowHeaders,
			"colHeaders":    upgradableColHeaders,
		})
		king.ForEachLevel(func(k king.King) {
			lv := k.Level()
			tableU.SetCell(0, lv-1, format.Int(king.ArenaTowerAtLevel(lv).Hitpoints()))
			tableU.SetCell(1, lv-1, format.Int(king.ArenaTowerAtLevel(lv).Damage()))
		})
		tableU.Print()
	}

	sep := ""
	typ.ForEach(func(t typ.Type) {
		// Card Type (Troops, Buildings, Spells)
		fmt.Printf(sep)
		fmt.Printf("## %s\n", t)
		fmt.Println()
		card.ForEachOfType(t, func(c card.Card) {

			// Header (Card Name)
			fmt.Printf("### %s\n", c.Name())
			fmt.Println()

			// Fixed Attribute Table
			{
				rowHeaders := []string{}
				contents := [][]string{}
				c.ForEachFixedAttribute(func(a attr.Fixed) {
					rowHeaders = append(rowHeaders, a.String())
					switch a {
					case attr.Rarity:
						contents = append(contents, []string{ilink(c.FormattedValue(a))})
					default:
						contents = append(contents, []string{c.FormattedValue(a)})
					}
				})
				table := NewTable(map[string]interface{}{
					"headerWidth":   fixedHeaderWidth,
					"contentsWidth": fixedContentsWidth,
					"rowHeaders":    rowHeaders,
					"colHeaders":    fixedColHeaders,
					"contents":      contents,
				})
				table.SetContentAlignment(LEFT)
				table.Print()
			}

			// Upgradable Attribute Table
			{
				rowHeaders := []string{}
				contents := [][]string{}
				c.ForEachUpgradableAttribute(func(a attr.Upgradable) {
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
