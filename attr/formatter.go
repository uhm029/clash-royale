package attr

import (
	"github.com/asukakenji/clash-royale/lib"
	"github.com/asukakenji/clash-royale/rng"

	"fmt"
)

func convertNumber(value interface{}) interface{} {
	switch value.(type) {
	case int:
		return value
	case float64:
		v := int(value.(float64))
		if float64(v) == value.(float64) {
			return v
		}
		return value
	default:
		panic("Unknown value type")
	}
}

func formatString(value interface{}) string {
	return fmt.Sprintf("%s", value)
}

func formatInt(value interface{}) string {
	// Note: interface{} is comparable with const
	if lib.X == value {
		return ""
	}
	return fmt.Sprintf("%d", value)
}

func formatFloat(value interface{}) string {
	// Note: interface{} is comparable with const
	if lib.X == value {
		return ""
	}
	number := convertNumber(value)
	switch number.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case float64:
		return fmt.Sprintf("%.1f", value)
	default:
		panic("Unknown value type")
	}
}

func formatElixir(value interface{}) string {
	if lib.X == value {
		return "?"
	}
	return formatInt(value)
}

func formatTime(value interface{}) string {
	number := convertNumber(value)
	switch number.(type) {
	case int:
		v := number.(int)
		if v < 60 {
			return fmt.Sprintf("%dsec", v)
		}
		return fmt.Sprintf("%dmin %dsec", v/60, v%60)
	case float64:
		return fmt.Sprintf("%.1fsec", value)
	default:
		panic("Unknown value type")
	}
}

func formatRange(value interface{}) string {
	switch value.(type) {
	case int:
		if value.(int) == rng.Melee {
			return "Melee"
		}
		return fmt.Sprintf("%d", value)
	case float64:
		return fmt.Sprintf("%.1f", value)
	default:
		panic("Unknown value type")
	}
}

func formatCount(value interface{}) string {
	return fmt.Sprintf("x %d", value.(int))
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
