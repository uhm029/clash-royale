package attribute

var (
	V1000R = []int{1000, 1100, 1210, 1330, 1460, 1600, 1760, 1930, 2120, 2330, 2560, 2810}
)

func generateHp(baseHp interface{}) []int {
	baseValue := baseHp.(int)
	values := make([]int, len(V1000R))
	for i, v := range V1000R {
		values[i] = v * baseValue / 1000
	}
	return values
}

func generateDam(baseDam interface{}) []int {
	return generateHp(baseDam)
}

func generateLv(baseLv interface{}) []int {
	baseValue := baseLv.(int)
	values := make([]int, 12)
	for i := range values {
		values[i] = baseValue + i
	}
	return values
}

func generateDur(baseDur interface{}) []float64 {
	baseValues := baseDur.([]float64)
	baseValue := baseValues[0]
	increment := baseValues[1]
	values := make([]float64, 12)
	for i := range values {
		values[i] = baseValue + float64(i)*increment
	}
	return values
}

/*
func generateDps(baseDam interface{}, hitSpeed float64) []int {
	values := generateDam(baseDam)
	for i, v := range values {
		values[i] = int(float64(v)/hitSpeed + 0.5)
	}
	return values
}
*/
