package lib

import (
	"fmt"
)

type Attribute interface {
	Attribute() // Tag
	String() string
}

type FixedAttribute struct {
	name        string
	FormatValue func(value interface{}) string
}

func (attr *FixedAttribute) Attribute() {
}

func (attr *FixedAttribute) String() string {
	return attr.name
}

type UpgradableAttribute struct {
	name         string
	FormatValues func(values interface{}) []string
}

func (attr *UpgradableAttribute) Attribute() {
}

func (attr *UpgradableAttribute) String() string {
	return attr.name
}

// --- Format Functions ---

func formatString(value interface{}) string {
	return fmt.Sprintf("%s", value)
}

func formatInt(value interface{}) string {
	// Note: interface{} is comparable with const
	if X == value {
		return ""
	}
	return fmt.Sprintf("%d", value)
}

func formatTime(value interface{}) string {
	switch value.(type) {
	case int:
		return fmt.Sprintf("%dsec", value)
	case float64:
		return fmt.Sprintf("%.1fsec", value)
	default:
		panic("Unknown value type")
	}
}

func formatInts(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = formatInt(v)
	}
	return strings
}

func formatTimes(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = formatTime(v)
	}
	return strings
}

var (
	NAME = &FixedAttribute{
		"Name",
		formatString,
	}
	ARENA = &FixedAttribute{
		"Arena",
		formatString,
	}
	RARITY = &FixedAttribute{
		"Rarity",
		formatString,
	}
	TYPE = &FixedAttribute{
		"Type",
		formatString,
	}
	DESC = &FixedAttribute{
		"Description",
		formatString,
	}
	COST = &FixedAttribute{
		"Elixir Cost",
		formatInt,
	}
	HP = &UpgradableAttribute{
		"Hitpoints",
		formatInts,
	}
	DPS = &UpgradableAttribute{
		"Damage per Second",
		formatInts,
	}
	DAM = &UpgradableAttribute{
		"Damage",
		formatInts,
	}
	ADAM = &UpgradableAttribute{
		"Area Damage",
		formatInts,
	}
	DDAM = &UpgradableAttribute{
		"Death Damage",
		formatInts,
	}
	SKE_LV = &UpgradableAttribute{
		"Skeleton Level",
		formatInts,
	}
	SGO_LV = &UpgradableAttribute{
		"Spear Goblin Level",
		formatInts,
	}
	SSPD = &FixedAttribute{
		"Spawn Speed",
		formatTime,
	}
	HSPD = &FixedAttribute{
		"Hit Speed",
		formatTime,
	}
	TGTS = &FixedAttribute{
		"Targets",
		formatString,
	}
	SPD = &FixedAttribute{
		"Speed",
		formatString,
	}
	RNG = &FixedAttribute{
		"Range",
		func(value interface{}) string {
			switch value.(type) {
			case int:
				if value.(int) == MELEE {
					return "Melee"
				}
				return fmt.Sprintf("%d", value)
			case float64:
				return fmt.Sprintf("%.1f", value)
			default:
				panic("Unknown value type")
			}
		},
	}
	DTIME = &FixedAttribute{
		"Deploy Time",
		formatTime,
	}
	COUNT = &FixedAttribute{
		"Count",
		func(value interface{}) string {
			return fmt.Sprintf("x %d", value.(int))
		},
	}
	CARDS_REQ = &UpgradableAttribute{
		"Cards Required",
		formatInts,
	}
	GOLD_REQ  = &UpgradableAttribute{
		"Gold Required",
		formatInts,
	}
	EXP_GAIN  = &UpgradableAttribute{
		"Experience Gained",
		formatInts,
	}
)

var ATTRIBUTES = [...]Attribute{
	NAME,
	ARENA,
	RARITY,
	TYPE,
	DESC,
	COST,
	HP,
	DPS,
	DAM,
	ADAM,
	DDAM,
	SKE_LV,
	SGO_LV,
	SSPD,
	HSPD,
	TGTS,
	SPD,
	RNG,
	DTIME,
	COUNT,
	CARDS_REQ,
	GOLD_REQ,
	EXP_GAIN,
}
