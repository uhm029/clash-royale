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
		"Damage per Second (L)",
		formatInts,
	}
	DPSH = &UpgradableAttribute{
		"Damage per Second (H)",
		formatInts,
	}
	CTDPS = &UpgradableAttribute{
		"Crown Tower Damage/sec",
		formatInts,
	}
	DAM = &UpgradableAttribute{
		"Damage",
		formatInts,
	}
	DAML = &UpgradableAttribute{
		"Damage (L)",
		formatInts,
	}
	DAMH = &UpgradableAttribute{
		"Damage (H)",
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
	CTDAM = &UpgradableAttribute{
		"Crown Tower Damage",
		formatInts,
	}
	GOB_LV = &UpgradableAttribute{
		"Goblin Level",
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
	DUR_F = &FixedAttribute{
		"Duration",
		formatTime,
	}
	DUR_U = &UpgradableAttribute{
		"Duration",
		formatTimes,
	}
	RAD = &FixedAttribute{
		"Radius",
		formatFloat,
	}
	COUNT = &FixedAttribute{
		"Count",
		formatCount,
	}
	GOB_COUNT = &FixedAttribute{
		"Goblin Count",
		formatCount,
	}
	MC_LV = &UpgradableAttribute{
		"Mirrored Common Level",
		formatInts,
	}
	MR_LV = &UpgradableAttribute{
		"Mirrored Rare Level",
		formatInts,
	}
	ME_LV = &UpgradableAttribute{
		"Mirrored Epic Level",
		formatInts,
	}
	ML_LV = &UpgradableAttribute{
		"Mirrored Ledendary Level",
		formatInts,
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
	CTDPS,
	DAM,
	DAML,
	DAMH,
	ADAM,
	DDAM,
	CTDAM,
	GOB_LV,
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
	DUR_F,
	DUR_U,
	RAD,
	COUNT,
	GOB_COUNT,
	MC_LV,
	MR_LV,
	ME_LV,
	ML_LV,
	CARDS_REQ,
	GOLD_REQ,
	EXP_GAIN,
}
