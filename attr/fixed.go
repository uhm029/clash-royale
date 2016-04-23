package attr

// Fixed
type Fixed int8

const (
	Name Fixed = iota
	Arena
	Rarity
	Type
	Desc
	Elixir
	SSpeed
	PSpeed
	HSpeed
	Targets
	Speed
	Range
	DTime
	LTime
	DurF
	Radius
	Count
	GobCount
)

func ForEachFixed(f func(Fixed)) {
	for i := range fixedAttributes {
		f(Fixed(i))
	}
}

func (a Fixed) Attribute() {
}

func (a Fixed) String() string {
	return fixedAttributes[a].name
}

func (a Fixed) Name() string {
	return fixedAttributes[a].name
}

func (a Fixed) FormatValue(value interface{}) string {
	return fixedAttributes[a].formatFunc(value)
}

/////////////
// Private //
/////////////

type fixedAttribute struct {
	name       string
	formatFunc func(value interface{}) string
}

var fixedAttributes = [...]*fixedAttribute{
	&fixedAttribute{"Name", formatString},
	&fixedAttribute{"Arena", formatString},
	&fixedAttribute{"Rarity", formatString},
	&fixedAttribute{"Type", formatString},
	&fixedAttribute{"Description", formatString},
	&fixedAttribute{"Elixir Cost", formatElixir},
	&fixedAttribute{"Spawn Speed", formatTime},
	&fixedAttribute{"Production Speed", formatTime},
	&fixedAttribute{"Hit Speed", formatTime},
	&fixedAttribute{"Targets", formatString},
	&fixedAttribute{"Speed", formatString},
	&fixedAttribute{"Range", formatRange},
	&fixedAttribute{"Deploy Time", formatTime},
	&fixedAttribute{"Lifetime", formatTime},
	&fixedAttribute{"Duration", formatTime},
	&fixedAttribute{"Radius", formatFloat},
	&fixedAttribute{"Count", formatCount},
	&fixedAttribute{"Goblin Count", formatCount},
}
