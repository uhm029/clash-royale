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
	DPSL = &UpgradableAttribute{
		"Damage per Second (Lowest)",
		formatInts,
	}
	DPSH = &UpgradableAttribute{
		"Damage per Second (Highest)",
		formatInts,
	}
	DAM = &UpgradableAttribute{
		"Damage",
		formatInts,
	}
	DAML = &UpgradableAttribute{
		"Damage (Lowest)",
		formatInts,
	}
	DAMH = &UpgradableAttribute{
		"Damage (Highest)",
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
	SGO_LV = &UpgradableAttribute{
		"Spear Goblin Level",
		formatInts,
	}
	SKE_LV = &UpgradableAttribute{
		"Skeleton Level",
		formatInts,
	}
	BAR_LV = &UpgradableAttribute{
		"Barbarian Level",
		formatInts,
	}
	SSPD = &FixedAttribute{
		"Spawn Speed",
		formatTime,
	}
	PSPD = &FixedAttribute{
		"Production Speed",
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
	LTIME = &FixedAttribute{
		"Lifetime",
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
	GOLD_REQ = &UpgradableAttribute{
		"Gold Required",
		formatInts,
	}
	EXP_GAIN = &UpgradableAttribute{
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
	DPSL,
	DPSH,
	DAM,
	DAML,
	DAMH,
	ADAM,
	DDAM,
	SGO_LV,
	SKE_LV,
	BAR_LV,
	SSPD,
	PSPD,
	HSPD,
	TGTS,
	SPD,
	RNG,
	DTIME,
	LTIME,
	COUNT,
	CARDS_REQ,
	GOLD_REQ,
	EXP_GAIN,
}
