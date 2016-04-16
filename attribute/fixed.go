package attribute

// Fixed
type Fixed struct {
	id          int
	name        string
	formatValue func(value interface{}) string
}

func (attr *Fixed) Attribute() {
}

func (attr *Fixed) Id() int {
	return attr.id
}

func (attr *Fixed) String() string {
	return attr.name
}

func (attr *Fixed) Name() string {
	return attr.name
}

func (attr *Fixed) FormatValue(value interface{}) string {
	return attr.formatValue(value)
}
