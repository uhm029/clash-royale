package arena

import (
	"fmt"
	"sort"
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
	arenas = []*Arena{}
)

// constructor
func newArena(id int, name string, trophies int) *Arena {
	a := &Arena{
		id,
		name,
		trophies,
	}
	arenas = append(arenas, a)
	return a
}

func ForEachArena(f func(*Arena)) {
	for _, a := range arenas {
		f(a)
	}
}

var (
	ARENA_0 = newArena(0, "Training Camp", -1)
	ARENA_1 = newArena(1, "Goblin Stadium", 0)
	ARENA_2 = newArena(2, "Bone Pit", 400)
	ARENA_3 = newArena(3, "Barbarian Bowl", 800)
	ARENA_4 = newArena(4, "P.E.K.K.A's Playhouse", 1100)
	ARENA_5 = newArena(5, "Spell Valley", 1400)
	ARENA_6 = newArena(6, "Builder's Workshop", 1700)
	ARENA_7 = newArena(7, "Royal Arena", 2000)
	ARENA_8 = newArena(8, "Legendary Arena", 3000)
)

// Initialization
type arenaSlice []*Arena

func (s arenaSlice) Len() int {
	return len(s)
}

func (s arenaSlice) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s arenaSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func init() {
	sort.Sort(arenaSlice(arenas))
}
