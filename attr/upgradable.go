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
	FspLV
	GolLV
	LavLV
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
	name       string
	formatFunc func(value interface{}) string
}

var upgradableAttributes = [...]*upgradableAttribute{
	&upgradableAttribute{"Hitpoints", formatInt},
	&upgradableAttribute{"Shield Hitpoints", formatInt},
	&upgradableAttribute{"Damage per Second", formatInt},
	&upgradableAttribute{"Damage per Second (L)", formatInt},
	&upgradableAttribute{"Damage per Second (H)", formatInt},
	&upgradableAttribute{"Crown Tower Damage/sec", formatInt},
	&upgradableAttribute{"Damage", formatInt},
	&upgradableAttribute{"Damage (L)", formatInt},
	&upgradableAttribute{"Damage (H)", formatInt},
	&upgradableAttribute{"Area Damage", formatInt},
	&upgradableAttribute{"Death Damage", formatInt},
	&upgradableAttribute{"Crown Tower Damage", formatInt},
	&upgradableAttribute{"Goblin Level", formatInt},
	&upgradableAttribute{"Spear Goblin Level", formatInt},
	&upgradableAttribute{"Skeleton Level", formatInt},
	&upgradableAttribute{"Barbarian Level", formatInt},
	&upgradableAttribute{"Fire Spirits Level", formatInt},
	&upgradableAttribute{"Golemite Level", formatInt},
	&upgradableAttribute{"Lava Pups Level", formatInt},
	&upgradableAttribute{"Duration", formatTime},
	&upgradableAttribute{"Mirrored Common Level", formatInt},
	&upgradableAttribute{"Mirrored Rare Level", formatInt},
	&upgradableAttribute{"Mirrored Epic Level", formatInt},
	&upgradableAttribute{"Mirrored Legendary Level", formatInt},
	&upgradableAttribute{"Cards Required", formatInt},
	&upgradableAttribute{"Gold Required", formatInt},
	&upgradableAttribute{"Experience Gained", formatInt},
}
