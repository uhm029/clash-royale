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

func (a Fixed) Id() int {
	return fixedAttributes[a].id
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
	id         int
	name       string
	formatFunc func(value interface{}) string
}

var fixedAttributes = []*fixedAttribute{
	&fixedAttribute{0, "Name", formatString},
	&fixedAttribute{1, "Arena", formatString},
	&fixedAttribute{2, "Rarity", formatString},
	&fixedAttribute{3, "Type", formatString},
	&fixedAttribute{4, "Description", formatString},
	&fixedAttribute{5, "Elixir Cost", formatInt},
	&fixedAttribute{500, "Spawn Speed", formatTime},
	&fixedAttribute{510, "Production Speed", formatTime},
	&fixedAttribute{520, "Hit Speed", formatTime},
	&fixedAttribute{600, "Targets", formatString},
	&fixedAttribute{700, "Speed", formatString},
	&fixedAttribute{800, "Range", formatRange},
	&fixedAttribute{900, "Deploy Time", formatTime},
	&fixedAttribute{910, "Lifetime", formatTime},
	&fixedAttribute{920, "Duration", formatTime},
	&fixedAttribute{1000, "Radius", formatFloat},
	&fixedAttribute{1100, "Count", formatCount},
	&fixedAttribute{1200, "Goblin Count", formatCount},
}
