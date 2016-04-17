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
	generateFunc    func(baseValue interface{}, max int) []interface{}
}

var generatedAttributes = []*generatedAttribute{
	&generatedAttribute{10100, HP, generateHp},
	&generatedAttribute{10110, SHP, generateHp},
	&generatedAttribute{10300, Dam, generateDam},
	&generatedAttribute{10310, DamL, generateDam},
	&generatedAttribute{10320, DamH, generateDam},
	&generatedAttribute{10330, ADam, generateDam},
	&generatedAttribute{10340, DDam, generateDam},
	&generatedAttribute{10400, GobLV, generateLv},
	&generatedAttribute{10410, SgoLV, generateLv},
	&generatedAttribute{10420, SkeLV, generateLv},
	&generatedAttribute{10430, BarLV, generateLv},
	&generatedAttribute{10921, DurU, generateDur},
	&generatedAttribute{11300, MCLV, generateLv},
	&generatedAttribute{11310, MRLV, generateLv},
	&generatedAttribute{11320, MELV, generateLv},
}
