package attr

// Upgradable
type Upgradable int8

const (
	HP Upgradable = iota
	SHP
	DPS
	DPSL
	DPSH
	CTDPS
	Dam
	DamL
	DamH
	ADam
	DDam
	CTDam
	GobLV
	SgoLV
	SkeLV
	BarLV
	DurU
	MCLV
	MRLV
	MELV
	MLLV
	CardsReq
	GoldReq
	ExpGain
)

func ForEachUpgradable(f func(Upgradable)) {
	for i := range upgradableAttributes {
		f(Upgradable(i))
	}
}

func (a Upgradable) Attribute() {
}

func (a Upgradable) Id() int {
	return upgradableAttributes[a].id
}

func (a Upgradable) String() string {
	return upgradableAttributes[a].name
}

func (a Upgradable) Name() string {
	return upgradableAttributes[a].name
}

func (a Upgradable) FormatValues(values []interface{}) []string {
	strings := make([]string, len(values))
	for i, v := range values {
		strings[i] = upgradableAttributes[a].formatFunc(v)
	}
	return strings
}

/////////////
// Private //
/////////////

type upgradableAttribute struct {
	id         int
	name       string
	formatFunc func(value interface{}) string
}

var upgradableAttributes = []*upgradableAttribute{
	&upgradableAttribute{100, "Hitpoints", formatInt},
	&upgradableAttribute{110, "Shield Hitpoints", formatInt},
	&upgradableAttribute{200, "Damage per Second", formatInt},
	&upgradableAttribute{210, "Damage per Second (L)", formatInt},
	&upgradableAttribute{220, "Damage per Second (H)", formatInt},
	&upgradableAttribute{230, "Crown Tower Damage/sec", formatInt},
	&upgradableAttribute{300, "Damage", formatInt},
	&upgradableAttribute{310, "Damage (L)", formatInt},
	&upgradableAttribute{320, "Damage (H)", formatInt},
	&upgradableAttribute{330, "Area Damage", formatInt},
	&upgradableAttribute{340, "Death Damage", formatInt},
	&upgradableAttribute{350, "Crown Tower Damage", formatInt},
	&upgradableAttribute{400, "Goblin Level", formatInt},
	&upgradableAttribute{410, "Spear Goblin Level", formatInt},
	&upgradableAttribute{420, "Skeleton Level", formatInt},
	&upgradableAttribute{430, "Barbarian Level", formatInt},
	&upgradableAttribute{921, "Duration", formatTime},
	&upgradableAttribute{1300, "Mirrored Common Level", formatInt},
	&upgradableAttribute{1310, "Mirrored Rare Level", formatInt},
	&upgradableAttribute{1320, "Mirrored Epic Level", formatInt},
	&upgradableAttribute{1330, "Mirrored Legendary Level", formatInt},
	&upgradableAttribute{9990, "Cards Required", formatInt},
	&upgradableAttribute{9991, "Gold Required", formatInt},
	&upgradableAttribute{9992, "Experience Gained", formatInt},
}
