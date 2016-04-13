package lib

// Type
type Type struct {
	id   int
	name string
}

// static
var typeCount = 0

// constructor
func newType(name string) *Type {
	id := typeCount
	typeCount++
	return &Type{
		id,
		name,
	}
}

func (t *Type) String() string {
	return t.name
}

var (
	TROOP    = newType("Troop")
	BUILDING = newType("Building")
	SPELL    = newType("Spell")
)

var TYPES = [...]*Type{
	TROOP,
	BUILDING,
	SPELL,
}
