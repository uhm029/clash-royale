package lib

import (
	"fmt"
)

type Arena struct {
	id       int
	name     string
	trophies int
}

func (a *Arena) String() string {
	return fmt.Sprintf("Arena %d: %s", a.id, a.name)
}

var (
	ARENA_0 = &Arena{0, "Training Camp", -1}
	ARENA_1 = &Arena{1, "Goblin Stadium", 0}
	ARENA_2 = &Arena{2, "Bone Pit", 400}
	ARENA_3 = &Arena{3, "Barbarian Bowl", 800}
	ARENA_4 = &Arena{4, "P.E.K.K.A's Playhouse", 1100}
	ARENA_5 = &Arena{5, "Spell Valley", 1400}
	ARENA_6 = &Arena{6, "Builder's Workshop", 1700}
	ARENA_7 = &Arena{7, "Royal Arena", 2000}
	ARENA_8 = &Arena{8, "Legendary Arena", 3000}
)

var ARENAS = [...]*Arena{
	ARENA_0,
	ARENA_1,
	ARENA_2,
	ARENA_3,
	ARENA_4,
	ARENA_5,
	ARENA_6,
	ARENA_7,
	ARENA_8,
}
