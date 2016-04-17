package internal

// FixedAttribute
type FixedAttribute struct {
	Id         int
	Name       string
	FormatFunc func(value interface{}) string
}

var FixedAttributes = []*FixedAttribute{
	&FixedAttribute{0, "Name", formatString},
	&FixedAttribute{1, "Arena", formatString},
	&FixedAttribute{2, "Rarity", formatString},
	&FixedAttribute{3, "Type", formatString},
	&FixedAttribute{4, "Description", formatString},
	&FixedAttribute{5, "Elixir Cost", formatInt},
	&FixedAttribute{500, "Spawn Speed", formatTime},
	&FixedAttribute{510, "Production Speed", formatTime},
	&FixedAttribute{520, "Hit Speed", formatTime},
	&FixedAttribute{600, "Targets", formatString},
	&FixedAttribute{700, "Speed", formatString},
	&FixedAttribute{800, "Range", formatRange},
	&FixedAttribute{900, "Deploy Time", formatTime},
	&FixedAttribute{910, "Lifetime", formatTime},
	&FixedAttribute{920, "Duration", formatTime},
	&FixedAttribute{1000, "Radius", formatFloat},
	&FixedAttribute{1100, "Count", formatCount},
	&FixedAttribute{1200, "Goblin Count", formatCount},
}
