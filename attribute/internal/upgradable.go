package internal

// UpgradableAttribute
type UpgradableAttribute struct {
	Id         int
	Name       string
	FormatFunc func(values interface{}) []string
}

var UpgradableAttributes = []*UpgradableAttribute{
	&UpgradableAttribute{100, "Hitpoints", formatInts},
	&UpgradableAttribute{110, "Shield Hitpoints", formatInts},
	&UpgradableAttribute{200, "Damage per Second", formatInts},
	&UpgradableAttribute{210, "Damage per Second (L)", formatInts},
	&UpgradableAttribute{220, "Damage per Second (H)", formatInts},
	&UpgradableAttribute{230, "Crown Tower Damage/sec", formatInts},
	&UpgradableAttribute{300, "Damage", formatInts},
	&UpgradableAttribute{310, "Damage (L)", formatInts},
	&UpgradableAttribute{320, "Damage (H)", formatInts},
	&UpgradableAttribute{330, "Area Damage", formatInts},
	&UpgradableAttribute{340, "Death Damage", formatInts},
	&UpgradableAttribute{350, "Crown Tower Damage", formatInts},
	&UpgradableAttribute{400, "Goblin Level", formatInts},
	&UpgradableAttribute{410, "Spear Goblin Level", formatInts},
	&UpgradableAttribute{420, "Skeleton Level", formatInts},
	&UpgradableAttribute{430, "Barbarian Level", formatInts},
	&UpgradableAttribute{921, "Duration", formatTimes},
	&UpgradableAttribute{1300, "Mirrored Common Level", formatInts},
	&UpgradableAttribute{1310, "Mirrored Rare Level", formatInts},
	&UpgradableAttribute{1320, "Mirrored Epic Level", formatInts},
	&UpgradableAttribute{1330, "Mirrored Legendary Level", formatInts},
	&UpgradableAttribute{9990, "Cards Required", formatInts},
	&UpgradableAttribute{9991, "Gold Required", formatInts},
	&UpgradableAttribute{9992, "Experience Gained", formatInts},
}

// This is copied from the "attribute" package to facilitate referral from "generated.go"
const (
	HP = iota
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
