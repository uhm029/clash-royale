package attr

// Attribute
type Attribute interface {
	Attribute() // Tag
}

func ForEach(f func(Attribute)) {
	for _, v := range attributes {
		f(v)
	}
}

// Helper
type BaseDuration struct {
	BaseValue int
	Increment float64
}

/////////////
// Private //
/////////////

var attributes = [...]Attribute{
	Name,
	Arena,
	Rarity,
	Type,
	Desc,
	Elixir,
	HP,
	SHP,
	HP,
	SHP,
	DPS,
	DPSL,
	DPSH,
	CTDPS,
	Dam,
	DamL,
	DamH,
	ADam,
	DDam,
	CTDam,
	GobLV,
	SgoLV,
	SkeLV,
	BarLV,
	SSpeed,
	PSpeed,
	HSpeed,
	Targets,
	Speed,
	Range,
	DTime,
	LTime,
	DurF,
	DurU,
	Radius,
	Count,
	GobCount,
	MCLV,
	MRLV,
	MELV,
	MLLV,
	CardsReq,
	GoldReq,
	ExpGain,
	BaseHP,
	BaseSHP,
	BaseDam,
	BaseDamL,
	BaseDamH,
	BaseADam,
	BaseDDam,
	BaseGobLV,
	BaseSgoLV,
	BaseSkeLV,
	BaseBarLV,
	BaseDurU,
	BaseMCLV,
	BaseMRLV,
	BaseMELV,
}
