package lib

import (
	"fmt"
)

type CardAttribute interface {
	CardAttribute() // Tag
	String() string
}

type FixedCardAttribute struct {
	name        string
	FormatValue func(value interface{}) string
}

func (attr *FixedCardAttribute) CardAttribute() {
}

func (attr *FixedCardAttribute) String() string {
	return attr.name
}

type UpgradableCardAttribute struct {
	name         string
	FormatValues func(values interface{}) []string
}

func (attr *UpgradableCardAttribute) CardAttribute() {
}

func (attr *UpgradableCardAttribute) String() string {
	return attr.name
}

// --- Format Functions ---

func FormatString(value interface{}) string {
	return value.(string)
}

func FormatInt(value interface{}) string {
	if X == value.(int) {
		return ""
	}
	return fmt.Sprintf("%d", value.(int))
}

func FormatTime(value interface{}) string {
	switch value.(type) {
	case int:
		return fmt.Sprintf("%dsec", value)
	case float64:
		return fmt.Sprintf("%.1fsec", value)
	default:
		panic("Unknown value type")
	}
}

func FormatInts(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = FormatInt(v)
	}
	return strings
}

func FormatTimes(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = FormatTime(v)
	}
	return strings
}

var (
	NAME = &FixedCardAttribute{
		"Name",
		FormatString,
	}
	ARENA = &FixedCardAttribute{
		"Arena",
		func(value interface{}) string {
			return value.(*Arena).String()
		},
	}
	RARITY = &FixedCardAttribute{
		"Rarity",
		func(value interface{}) string {
			return value.(*Rarity).String()
		},
	}
	TYPE = &FixedCardAttribute{
		"Type",
		func(value interface{}) string {
			return string(value.(Type))
		},
	}
	DESC = &FixedCardAttribute{
		"Description",
		FormatString,
	}
	COST = &FixedCardAttribute{
		"Elixir Cost",
		FormatInt,
	}
	HP = &UpgradableCardAttribute{
		"Hitpoints",
		FormatInts,
	}
	DPS = &UpgradableCardAttribute{
		"Damage per Second",
		FormatInts,
	}
	DAM = &UpgradableCardAttribute{
		"Damage",
		FormatInts,
	}
	ADAM = &UpgradableCardAttribute{
		"Area Damage",
		FormatInts,
	}
	DDAM = &UpgradableCardAttribute{
		"Death Damage",
		FormatInts,
	}
	SKE_LV = &UpgradableCardAttribute{
		"Skeleton Level",
		FormatInts,
	}
	SGO_LV = &UpgradableCardAttribute{
		"Spear Goblin Level",
		FormatInts,
	}
	SSPD = &FixedCardAttribute{
		"Spawn Speed",
		FormatTime,
	}
	HSPD = &FixedCardAttribute{
		"Hit Speed",
		FormatTime,
	}
	TGTS = &FixedCardAttribute{
		"Targets",
		func(value interface{}) string {
			return string(value.(Targets))
		},
	}
	SPD = &FixedCardAttribute{
		"Speed",
		func(value interface{}) string {
			return string(value.(Speed))
		},
	}
	RNG = &FixedCardAttribute{
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
	DTIME = &FixedCardAttribute{
		"Deploy Time",
		FormatTime,
	}
	COUNT = &FixedCardAttribute{
		"Count",
		func(value interface{}) string {
			return fmt.Sprintf("x %d", value.(int))
		},
	}
)

var CARD_ATTRIBUTES = [...]CardAttribute{
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
}
