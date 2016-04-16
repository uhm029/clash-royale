package attribute

import (
	"github.com/asukakenji/clash-royale/attribute/internal"
)

// Fixed
type Fixed int8

const (
	Name Fixed = iota
	Arena
	Rarity
	Type
	Desc
	Elixir
	SSpeed
	PSpeed
	HSpeed
	Targets
	Speed
	Range
	DTime
	LTime
	DurF
	Radius
	Count
	GobCount
)

func ForEachFixed(f func(Fixed)) {
	for i := range internal.FixedAttributes {
		f(Fixed(i))
	}
}

func (a Fixed) Attribute() {
}

func (a Fixed) Id() int {
	return internal.FixedAttributes[a].Id
}

func (a Fixed) String() string {
	return internal.FixedAttributes[a].Name
}

func (a Fixed) Name() string {
	return internal.FixedAttributes[a].Name
}

func (a Fixed) FormatValue(value interface{}) string {
	return internal.FixedAttributes[a].FormatFunc(value)
}
