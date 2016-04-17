package attr

import (
	"github.com/asukakenji/clash-royale/attr/internal"
)

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
	for i := range internal.UpgradableAttributes {
		f(Upgradable(i))
	}
}

func (a Upgradable) Attribute() {
}

func (a Upgradable) Id() int {
	return internal.UpgradableAttributes[a].Id
}

func (a Upgradable) String() string {
	return internal.UpgradableAttributes[a].Name
}

func (a Upgradable) Name() string {
	return internal.UpgradableAttributes[a].Name
}

func (a Upgradable) FormatValues(values interface{}) []string {
	return internal.UpgradableAttributes[a].FormatFunc(values)
}
