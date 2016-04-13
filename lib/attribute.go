package lib

// Attribute
type Attribute interface {
	Attribute() // Tag
	String() string
}

// FixedAttribute
type FixedAttribute struct {
	id          int
	name        string
	formatValue func(value interface{}) string
}

// static
var fixedAttributeCount = 0

// constructor
func newFixedAttribute(name string, formatValue func(value interface{}) string) *FixedAttribute {
	id := fixedAttributeCount
	fixedAttributeCount++
	return &FixedAttribute{
		id,
		name,
		formatValue,
	}
}

func (attr *FixedAttribute) Attribute() {
}

func (attr *FixedAttribute) String() string {
	return attr.name
}

// UpgradableAttribute
type UpgradableAttribute struct {
	id           int
	name         string
	formatValues func(values interface{}) []string
}

// static
var upgradableAttributeCount = 0

// constructor
func newUpgradableAttribute(name string, formatValues func(values interface{}) []string) *UpgradableAttribute {
	id := upgradableAttributeCount
	upgradableAttributeCount++
	return &UpgradableAttribute{
		id,
		name,
		formatValues,
	}
}

func (attr *UpgradableAttribute) Attribute() {
}

func (attr *UpgradableAttribute) String() string {
	return attr.name
}

var (
	NAME      = newFixedAttribute("Name", formatString)
	ARENA     = newFixedAttribute("Arena", formatString)
	RARITY    = newFixedAttribute("Rarity", formatString)
	TYPE      = newFixedAttribute("Type", formatString)
	DESC      = newFixedAttribute("Description", formatString)
	COST      = newFixedAttribute("Elixir Cost", formatInt)
	HP        = newUpgradableAttribute("Hitpoints", formatInts)
	SHP       = newUpgradableAttribute("Shield Hitpoints", formatInts)
	DPS       = newUpgradableAttribute("Damage per Second", formatInts)
	DPSL      = newUpgradableAttribute("Damage per Second (L)", formatInts)
	DPSH      = newUpgradableAttribute("Damage per Second (H)", formatInts)
	CTDPS     = newUpgradableAttribute("Crown Tower Damage/sec", formatInts)
	DAM       = newUpgradableAttribute("Damage", formatInts)
	DAML      = newUpgradableAttribute("Damage (L)", formatInts)
	DAMH      = newUpgradableAttribute("Damage (H)", formatInts)
	ADAM      = newUpgradableAttribute("Area Damage", formatInts)
	DDAM      = newUpgradableAttribute("Death Damage", formatInts)
	CTDAM     = newUpgradableAttribute("Crown Tower Damage", formatInts)
	GOB_LV    = newUpgradableAttribute("Goblin Level", formatInts)
	SGO_LV    = newUpgradableAttribute("Spear Goblin Level", formatInts)
	SKE_LV    = newUpgradableAttribute("Skeleton Level", formatInts)
	BAR_LV    = newUpgradableAttribute("Barbarian Level", formatInts)
	SSPD      = newFixedAttribute("Spawn Speed", formatTime)
	PSPD      = newFixedAttribute("Production Speed", formatTime)
	HSPD      = newFixedAttribute("Hit Speed", formatTime)
	TGTS      = newFixedAttribute("Targets", formatString)
	SPD       = newFixedAttribute("Speed", formatString)
	RNG       = newFixedAttribute("Range", formatRange)
	DTIME     = newFixedAttribute("Deploy Time", formatTime)
	LTIME     = newFixedAttribute("Lifetime", formatTime)
	DUR_F     = newFixedAttribute("Duration", formatTime)
	DUR_U     = newUpgradableAttribute("Duration", formatTimes)
	RAD       = newFixedAttribute("Radius", formatFloat)
	COUNT     = newFixedAttribute("Count", formatCount)
	GOB_COUNT = newFixedAttribute("Goblin Count", formatCount)
	MC_LV     = newUpgradableAttribute("Mirrored Common Level", formatInts)
	MR_LV     = newUpgradableAttribute("Mirrored Rare Level", formatInts)
	ME_LV     = newUpgradableAttribute("Mirrored Epic Level", formatInts)
	ML_LV     = newUpgradableAttribute("Mirrored Legendary Level", formatInts)
	CARDS_REQ = newUpgradableAttribute("Cards Required", formatInts)
	GOLD_REQ  = newUpgradableAttribute("Gold Required", formatInts)
	EXP_GAIN  = newUpgradableAttribute("Experience Gained", formatInts)
)

var ATTRIBUTES = [...]Attribute{
	NAME,
	ARENA,
	RARITY,
	TYPE,
	DESC,
	COST,
	HP,
	SHP,
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
