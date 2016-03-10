package lib

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
		formatRange,
	}
	DTIME = &FixedAttribute{
		"Deploy Time",
		formatTime,
	}
	COUNT = &FixedAttribute{
		"Count",
		formatCount,
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
