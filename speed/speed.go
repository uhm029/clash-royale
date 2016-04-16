package speed

import (
	"github.com/asukakenji/clash-royale/speed/internal"
)

// Speed
type Speed int8

const (
	Slow Speed = iota
	Medium
	Fast
	VeryFast
)

func ForEach(f func(Speed)) {
	for i := range internal.Speeds {
		f(Speed(i))
	}
}

func (s Speed) Id() int {
	return internal.Speeds[s].Id
}

func (s Speed) String() string {
	return internal.Speeds[s].Name
}

func (s Speed) Name() string {
	return internal.Speeds[s].Name
}
