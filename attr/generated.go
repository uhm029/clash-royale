package attr

// Generated
type Generated int8

const (
	BaseHP Generated = iota
	BaseSHP
	BaseDam
	BaseDamL
	BaseDamH
	BaseADam
	BaseDDam
	BaseGobLV
	BaseSgoLV
	BaseSkeLV
	BaseBarLV
	BaseDurU
	BaseMCLV
	BaseMRLV
	BaseMELV
)

func ForEachGenerated(f func(Generated)) {
	for i := range generatedAttributes {
		f(Generated(i))
	}
}

func (a Generated) Attribute() {
}

func (a Generated) Id() int {
	return generatedAttributes[a].id
}

func (a Generated) String() string {
	return a.TargetAttribute().String()
}

func (a Generated) TargetAttribute() Upgradable {
	return generatedAttributes[a].targetAttribute
}

func (a Generated) GenerateValues(baseValue interface{}, max int) interface{} {
	return generatedAttributes[a].generateFunc(baseValue, max)
}

/////////////
// Private //
/////////////

type generatedAttribute struct {
	id              int
	targetAttribute Upgradable
	generateFunc    func(baseValue interface{}, max int) interface{}
}

func i2si(f func(int, int) []int) func(interface{}, int) interface{} {
	return func(value interface{}, max int) interface{} {
		return f(value.(int), max)
	}
}

func xbd2so(f func(BaseDuration, int) []interface{}) func(interface{}, int) interface{} {
	return func(value interface{}, max int) interface{} {
		return f(value.(BaseDuration), max)
	}
}

var generatedAttributes = []*generatedAttribute{
	&generatedAttribute{10100, HP, i2si(generateHp)},
	&generatedAttribute{10110, SHP, i2si(generateHp)},
	&generatedAttribute{10300, Dam, i2si(generateDam)},
	&generatedAttribute{10310, DamL, i2si(generateDam)},
	&generatedAttribute{10320, DamH, i2si(generateDam)},
	&generatedAttribute{10330, ADam, i2si(generateDam)},
	&generatedAttribute{10340, DDam, i2si(generateDam)},
	&generatedAttribute{10400, GobLV, i2si(generateLv)},
	&generatedAttribute{10410, SgoLV, i2si(generateLv)},
	&generatedAttribute{10420, SkeLV, i2si(generateLv)},
	&generatedAttribute{10430, BarLV, i2si(generateLv)},
	&generatedAttribute{10921, DurU, xbd2so(generateDur)},
	&generatedAttribute{11300, MCLV, i2si(generateLv)},
	&generatedAttribute{11310, MRLV, i2si(generateLv)},
	&generatedAttribute{11320, MELV, i2si(generateLv)},
}
