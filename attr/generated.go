package attribute

import (
	"github.com/asukakenji/clash-royale/attribute/internal"
)

// Generated
type Generated int8

const (
	BaseHP Generated = iota
	BaseSHP
	BaseDam
	BaseDamL
	BaseDamH
	BaseADam
	BaseDDam
	BaseGobLV
	BaseSgoLV
	BaseSkeLV
	BaseBarLV
	BaseMCLV
	BaseMRLV
	BaseMELV
)

func ForEachGenerated(f func(Generated)) {
	for i := range internal.GeneratedAttributes {
		f(Generated(i))
	}
}

func (a Generated) Attribute() {
}

func (a Generated) Id() int {
	return internal.GeneratedAttributes[a].Id
}

func (a Generated) String() string {
	return a.TargetAttribute().String()
}

func (a Generated) TargetAttribute() Upgradable {
	return Upgradable(internal.GeneratedAttributes[a].AttributeIndex)
}

func (a Generated) GenerateValues(baseValue interface{}) interface{} {
	return internal.GeneratedAttributes[a].GenerateFunc(baseValue)
}
