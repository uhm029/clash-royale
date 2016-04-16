package arena

import (
	"github.com/asukakenji/clash-royale/arena/internal"
)

// Arena
type Arena byte

const (
	Arena0 Arena = iota
	Arena1
	Arena2
	Arena3
	Arena4
	Arena5
	Arena6
	Arena7
	Arena8
)

func ForEach(f func(Arena)) {
	for i := range internal.Arenas {
		f(Arena(i))
	}
}

func (a Arena) Id() int {
	return internal.Arenas[a].Id()
}

func (a Arena) String() string {
	return internal.Arenas[a].String()
}

func (a Arena) Name() string {
	return internal.Arenas[a].Name()
}

func (a Arena) Trophies() int {
	return internal.Arenas[a].Trophies()
}
