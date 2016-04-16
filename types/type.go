package types

import (
	"sort"
)

// Type
type Type struct {
	id   int
	name string
}

func (t *Type) Id() int {
	return t.id
}

func (t *Type) String() string {
	return t.name
}

func (t *Type) GetName() string {
	return t.name
}

// static
var (
	types = []*Type{}
)

// constructor
func newType(id int, name string) *Type {
	t := &Type{
		id,
		name,
	}
	types = append(types, t)
	return t
}

func ForEachType(f func(*Type)) {
	for _, t := range types {
		f(t)
	}
}

var (
	TROOP    = newType(0, "Troop")
	BUILDING = newType(1, "Building")
	SPELL    = newType(2, "Spell")
)

// Initialization
type typeSlice []*Type

func (s typeSlice) Len() int {
	return len(s)
}

func (s typeSlice) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s typeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func init() {
	sort.Sort(typeSlice(types))
}
