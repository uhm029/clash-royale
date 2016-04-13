package lib

import (
	"fmt"
)

// Arena
type Arena struct {
	id       int
	name     string
	trophies int
}

// static
var arenaCount = 0

// constructor
func newArena(name string, trophies int) *Arena {
	id := arenaCount
	arenaCount++
	return &Arena{
		id,
		name,
		trophies,
	}
}

func (a *Arena) String() string {
	return fmt.Sprintf("Arena %d: %s", a.id, a.name)
}

var (
	ARENA_0 = newArena("Training Camp", -1)
	ARENA_1 = newArena("Goblin Stadium", 0)
	ARENA_2 = newArena("Bone Pit", 400)
	ARENA_3 = newArena("Barbarian Bowl", 800)
	ARENA_4 = newArena("P.E.K.K.A's Playhouse", 1100)
	ARENA_5 = newArena("Spell Valley", 1400)
	ARENA_6 = newArena("Builder's Workshop", 1700)
	ARENA_7 = newArena("Royal Arena", 2000)
	ARENA_8 = newArena("Legendary Arena", 3000)
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
