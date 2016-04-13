package lib

// Targets
type Targets struct {
	id   int
	name string
}

// static
var targetsCount = 0

// constructor
func newTargets(name string) *Targets {
	id := targetsCount
	targetsCount++
	return &Targets{
		id,
		name,
	}
}

func (t *Targets) String() string {
	return t.name
}

var (
	GROUND         = newTargets("Ground")
	AIR_AND_GROUND = newTargets("Air & Ground")
	BUILDINGS      = newTargets("Buildings")
)

var TARGETSES = [...]*Targets{
	GROUND,
	AIR_AND_GROUND,
	BUILDINGS,
}
