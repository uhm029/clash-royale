package lib

// Speed
type Speed struct {
	id   int
	name string
}

// static
var speedCount = 0

// constructor
func newSpeed(name string) *Speed {
	id := speedCount
	speedCount++
	return &Speed{
		id,
		name,
	}
}

func (s *Speed) String() string {
	return s.name
}

var (
	SLOW      = newSpeed("Slow")
	MEDIUM    = newSpeed("Medium")
	FAST      = newSpeed("Fast")
	VERY_FAST = newSpeed("Very Fast")
)

var SPEEDS = [...]*Speed{
	SLOW,
	MEDIUM,
	FAST,
	VERY_FAST,
}
