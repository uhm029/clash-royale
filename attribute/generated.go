package attribute

// Generated
type Generated struct {
	id             int
	uattr          *Upgradable
	generateValues func(baseValue interface{}) []int
}

func (attr *Generated) Attribute() {
}

func (attr *Generated) Id() int {
	return attr.id
}

func (attr *Generated) String() string {
	return attr.uattr.String()
}

func (attr *Generated) Upgradable() *Upgradable {
	return attr.uattr
}

func (attr *Generated) GenerateValues(baseValue interface{}) []int {
	return attr.generateValues(baseValue)
}
