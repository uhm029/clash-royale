package lib

type Speed string

const (
	SLOW      = Speed("Slow")
	MEDIUM    = Speed("Medium")
	FAST      = Speed("Fast")
	VERY_FAST = Speed("Very Fast")
)

var SPEEDS = [...]Speed{
	SLOW,
	MEDIUM,
	FAST,
	VERY_FAST,
}
