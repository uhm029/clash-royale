package Type

import (
	"github.com/asukakenji/clash-royale/Type/internal"
)

// Type
type Type int8

const (
	Troop Type = iota
	Building
	Spell
)

func ForEach(f func(Type)) {
	for i := range internal.Types {
		f(Type(i))
	}
}

func (t Type) Id() int {
	return internal.Types[t].Id
}

func (t Type) String() string {
	return internal.Types[t].Name
}

func (t Type) Name() string {
	return internal.Types[t].Name
}
