package lib

import (
	"sort"
)

// Targets
type Targets struct {
	id   int
	name string
}

func (t *Targets) Id() int {
	return t.id
}

func (t *Targets) String() string {
	return t.name
}

func (t *Targets) GetName() string {
	return t.name
}

// static
var (
	targetses = []*Targets{}
)

// constructor
func newTargets(id int, name string) *Targets {
	t := &Targets{
		id,
		name,
	}
	targetses = append(targetses, t)
	return t
}

func ForEachTargets(f func(*Targets)) {
	for _, t := range targetses {
		f(t)
	}
}

var (
	GROUND         = newTargets(0, "Ground")
	AIR_AND_GROUND = newTargets(1, "Air & Ground")
	BUILDINGS      = newTargets(2, "Buildings")
)

// Initialization
type targetsSlice []*Targets

func (s targetsSlice) Len() int {
	return len(s)
}

func (s targetsSlice) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s targetsSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func init() {
	sort.Sort(targetsSlice(targetses))
}
