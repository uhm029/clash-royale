package internal

import (
	"fmt"
)

// Arena
type Arena struct {
	id       int
	name     string
	trophies int
}

func (a *Arena) Id() int {
	return a.id
}

func (a *Arena) String() string {
	return fmt.Sprintf("Arena %d: %s", a.id, a.name)
}

func (a *Arena) Name() string {
	return a.name
}

func (a *Arena) Trophies() int {
	return a.trophies
}

// static
var (
	Arenas = []*Arena{
		&Arena{0, "Training Camp", -1},
		&Arena{1, "Goblin Stadium", 0},
		&Arena{2, "Bone Pit", 400},
		&Arena{3, "Barbarian Bowl", 800},
		&Arena{4, "P.E.K.K.A's Playhouse", 1100},
		&Arena{5, "Spell Valley", 1400},
		&Arena{6, "Builder's Workshop", 1700},
		&Arena{7, "Royal Arena", 2000},
		&Arena{8, "Legendary Arena", 3000},
	}
)
