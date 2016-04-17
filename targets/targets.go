package targets

// Targets
type Targets int8

const (
	Ground Targets = iota
	AirAndGround
	Buildings
)

func ForEach(f func(Targets)) {
	for i := range targetses {
		f(Targets(i))
	}
}

func (t Targets) Id() int {
	return targetses[t].id
}

func (t Targets) String() string {
	return targetses[t].name
}

func (t Targets) Name() string {
	return targetses[t].name
}

/////////////
// Private //
/////////////

type targets struct {
	id   int
	name string
}

var targetses = []*targets{
	&targets{0, "Ground"},
	&targets{1, "Air & Ground"},
	&targets{2, "Buildings"},
}
