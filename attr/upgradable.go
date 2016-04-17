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

func (a Upgradable) FormatValues(values interface{}) []string {
	return upgradableAttributes[a].formatFunc(values)
}

/////////////
// Private //
/////////////

type upgradableAttribute struct {
	id         int
	name       string
	formatFunc func(values interface{}) []string
}

var upgradableAttributes = []*upgradableAttribute{
	&upgradableAttribute{100, "Hitpoints", formatInts},
	&upgradableAttribute{110, "Shield Hitpoints", formatInts},
	&upgradableAttribute{200, "Damage per Second", formatInts},
	&upgradableAttribute{210, "Damage per Second (L)", formatInts},
	&upgradableAttribute{220, "Damage per Second (H)", formatInts},
	&upgradableAttribute{230, "Crown Tower Damage/sec", formatInts},
	&upgradableAttribute{300, "Damage", formatInts},
	&upgradableAttribute{310, "Damage (L)", formatInts},
	&upgradableAttribute{320, "Damage (H)", formatInts},
	&upgradableAttribute{330, "Area Damage", formatInts},
	&upgradableAttribute{340, "Death Damage", formatInts},
	&upgradableAttribute{350, "Crown Tower Damage", formatInts},
	&upgradableAttribute{400, "Goblin Level", formatInts},
	&upgradableAttribute{410, "Spear Goblin Level", formatInts},
	&upgradableAttribute{420, "Skeleton Level", formatInts},
	&upgradableAttribute{430, "Barbarian Level", formatInts},
	&upgradableAttribute{921, "Duration", formatTimes},
	&upgradableAttribute{1300, "Mirrored Common Level", formatInts},
	&upgradableAttribute{1310, "Mirrored Rare Level", formatInts},
	&upgradableAttribute{1320, "Mirrored Epic Level", formatInts},
	&upgradableAttribute{1330, "Mirrored Legendary Level", formatInts},
	&upgradableAttribute{9990, "Cards Required", formatInts},
	&upgradableAttribute{9991, "Gold Required", formatInts},
	&upgradableAttribute{9992, "Experience Gained", formatInts},
}
