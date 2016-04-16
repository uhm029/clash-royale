package internal

// Targets
type Targets struct {
	Id   int
	Name string
}

var (
	Targetses = []*Targets{
		&Targets{0, "Ground"},
		&Targets{1, "Air & Ground"},
		&Targets{2, "Buildings"},
	}
)
