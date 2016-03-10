package lib

import (
	"fmt"
)

func formatString(value interface{}) string {
	return fmt.Sprintf("%s", value)
}

func formatInt(value interface{}) string {
	// Note: interface{} is comparable with const
	if X == value {
		return ""
	}
	return fmt.Sprintf("%d", value)
}

func formatTime(value interface{}) string {
	switch value.(type) {
	case int:
		v := value.(int)
		if value.(int) < 60 {
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
		if value.(int) == MELEE {
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
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = formatTime(v)
	}
	return strings
}
