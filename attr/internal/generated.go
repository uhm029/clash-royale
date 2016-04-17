package internal

// GeneratedAttribute
type GeneratedAttribute struct {
	Id             int
	AttributeIndex int
	GenerateFunc   func(baseValue interface{}) []int
}

var GeneratedAttributes = []*GeneratedAttribute{
	&GeneratedAttribute{10100, HP, generateHp},
	&GeneratedAttribute{10110, SHP, generateHp},
	&GeneratedAttribute{10300, Dam, generateDam},
	&GeneratedAttribute{10310, DamL, generateDam},
	&GeneratedAttribute{10320, DamH, generateDam},
	&GeneratedAttribute{10330, ADam, generateDam},
	&GeneratedAttribute{10340, DDam, generateDam},
	&GeneratedAttribute{10400, GobLV, generateLv},
	&GeneratedAttribute{10410, SgoLV, generateLv},
	&GeneratedAttribute{10420, SkeLV, generateLv},
	&GeneratedAttribute{10430, BarLV, generateLv},
	&GeneratedAttribute{11300, MCLV, generateLv},
	&GeneratedAttribute{11310, MRLV, generateLv},
	&GeneratedAttribute{11320, MELV, generateLv},
}
