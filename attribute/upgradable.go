package attribute

// Upgradable
type Upgradable struct {
	id           int
	name         string
	formatValues func(values interface{}) []string
}

func (attr *Upgradable) Attribute() {
}

func (attr *Upgradable) Id() int {
	return attr.id
}

func (attr *Upgradable) String() string {
	return attr.name
}

func (attr *Upgradable) Name() string {
	return attr.name
}

func (attr *Upgradable) FormatValues(values interface{}) []string {
	return attr.formatValues(values)
}
