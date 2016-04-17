package Type

// Type
type Type int8

const (
	Troop Type = iota
	Building
	Spell
)

func ForEach(f func(Type)) {
	for i := range types {
		f(Type(i))
	}
}

func (t Type) Id() int {
	return types[t].id
}

func (t Type) String() string {
	return types[t].name
}

func (t Type) Name() string {
	return types[t].name
}

/////////////
// Private //
/////////////

type _type struct {
	id   int
	name string
}

var types = []*_type{
	&_type{0, "Troop"},
	&_type{1, "Building"},
	&_type{2, "Spell"},
}
