package targets

import (
	"github.com/asukakenji/clash-royale/targets/internal"
)

type Targets int8

const (
	Ground Targets = iota
	AirAndGround
	Buildings
)

func ForEach(f func(Targets)) {
	for i := range internal.Targetses {
		f(Targets(i))
	}
}

func (t Targets) Id() int {
	return internal.Targetses[t].Id
}

func (t Targets) String() string {
	return internal.Targetses[t].Name
}

func (t Targets) Name() string {
	return internal.Targetses[t].Name
}
