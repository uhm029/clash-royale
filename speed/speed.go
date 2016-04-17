package speed

// Speed
type Speed int8

const (
	Slow Speed = iota
	Medium
	Fast
	VeryFast
)

func ForEach(f func(Speed)) {
	for i := range speeds {
		f(Speed(i))
	}
}

func (s Speed) Id() int {
	return speeds[s].id
}

func (s Speed) String() string {
	return speeds[s].name
}

func (s Speed) Name() string {
	return speeds[s].name
}

/////////////
// Private //
/////////////

type speed struct {
	id   int
	name string
}

var speeds = []*speed{
	&speed{0, "Slow"},
	&speed{1, "Medium"},
	&speed{2, "Fast"},
	&speed{3, "Very Fast"},
}
