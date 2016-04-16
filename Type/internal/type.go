package internal

// Type
type Type struct {
	Id   int
	Name string
}

var (
	Types = []*Type{
		&Type{0, "Troop"},
		&Type{1, "Building"},
		&Type{2, "Spell"},
	}
)
