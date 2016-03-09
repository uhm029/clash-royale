package lib

type Targets string

const (
	GROUND         = Targets("Ground")
	AIR_AND_GROUND = Targets("Air & Ground")
	BUILDINGS      = Targets("Buildings")
)

var TARGETSES = [...]Targets{
	GROUND,
	AIR_AND_GROUND,
	BUILDINGS,
}
