package lib

type Type string

const (
	TROOP    = Type("Troop")
	BUILDING = Type("Building")
	SPELL    = Type("Spell")
)

var TYPES = [...]Type{
	TROOP,
	BUILDING,
	SPELL,
}
