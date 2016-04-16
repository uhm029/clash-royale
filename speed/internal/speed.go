package internal

// Speed
type Speed struct {
	Id   int
	Name string
}

var (
	Speeds = []*Speed{
		&Speed{0, "Slow"},
		&Speed{1, "Medium"},
		&Speed{2, "Fast"},
		&Speed{3, "Very Fast"},
	}
)
