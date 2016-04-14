package lib

import (
	"sort"
)

// Speed
type Speed struct {
	id   int
	name string
}

func (s *Speed) Id() int {
	return s.id
}

func (s *Speed) String() string {
	return s.name
}

func (s *Speed) GetName() string {
	return s.name
}

// static
var (
	speedCount = 0
	speeds     = []*Speed{}
)

// constructor
func newSpeed(id int, name string) *Speed {
	s := &Speed{
		id,
		name,
	}
	speeds = append(speeds, s)
	return s
}

func ForEachSpeed(f func(*Speed)) {
	for _, s := range speeds {
		f(s)
	}
}

var (
	SLOW      = newSpeed(0, "Slow")
	MEDIUM    = newSpeed(1, "Medium")
	FAST      = newSpeed(2, "Fast")
	VERY_FAST = newSpeed(3, "Very Fast")
)

// Initialization
type speedSlice []*Speed

func (s speedSlice) Len() int {
	return len(s)
}

func (s speedSlice) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s speedSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func init() {
	sort.Sort(speedSlice(speeds))
}
