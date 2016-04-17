package attr

import (
	"github.com/asukakenji/clash-royale/common"
	"github.com/asukakenji/clash-royale/rng"

	"fmt"
)

// convertNumber converts a number into an appropriate type.
// An int or a float64 with fractional part is not converted.
// A float64 without fractional part is converted to an int.
func convertNumber(value interface{}) interface{} {
	switch v := value.(type) {
	case int:
		return v
	case float64:
		if i := int(v); float64(i) == v {
			return i
		}
		return v
	default:
		panic("Unknown value type")
	}
}

func formatString(value interface{}) string {
	return fmt.Sprintf("%s", value)
}

func formatInt(value interface{}) string {
	// Note: interface{} is comparable with const
	if common.X == value {
		return ""
	}
	return fmt.Sprintf("%d", value)
}

func formatFloat(value interface{}) string {
	number := convertNumber(value)
	switch v := number.(type) {
	case int:
		return formatInt(v)
	case float64:
		return fmt.Sprintf("%.1f", v)
	default:
		panic("Unknown value type")
	}
}

func formatElixir(value interface{}) string {
	if common.X == value {
		return "?"
	}
	return formatInt(value)
}

func formatTime(value interface{}) string {
	number := convertNumber(value)
	switch v := number.(type) {
	case int:
		if v < 60 {
			return fmt.Sprintf("%dsec", v)
		}
		return fmt.Sprintf("%dmin %dsec", v/60, v%60)
	case float64:
		return fmt.Sprintf("%.1fsec", v)
	default:
		panic("Unknown value type")
	}
}

func formatRange(value interface{}) string {
	if rng.Melee == value {
		return "Melee"
	}
	return formatFloat(value)
}

func formatCount(value interface{}) string {
	return fmt.Sprintf("x %d", value)
}

func formatInts(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = formatInt(v)
	}
	return strings
}

func formatTimes(values interface{}) []string {
	if ints, ok := values.([]int); ok {
		strings := make([]string, len(ints))
		for i, v := range ints {
			strings[i] = formatTime(v)
		}
		return strings
	} else if floats, ok := values.([]float64); ok {
		strings := make([]string, len(floats))
		for i, v := range floats {
			strings[i] = formatTime(v)
		}
		return strings
	} else {
		panic("Unknown value type")
	}
}
