package attr

var refValues = []int{1000, 1100, 1210, 1330, 1460, 1600, 1760, 1930, 2120, 2330, 2560, 2810}

const maxLevel = 12

func generateHp(baseHp, max int) []int {
	values := make([]int, maxLevel)
	for i := range values {
		values[i] = refValues[i] * baseHp / refValues[0]
	}
	return values[0:max:max]
}

func generateDam(baseDam, max int) []int {
	return generateHp(baseDam, max)
}

func generateLv(baseLv, max int) []int {
	values := make([]int, maxLevel)
	for i := range values {
		values[i] = baseLv + i
	}
	return values[0:max:max]
}

func generateDur(baseDur BaseDuration, max int) []interface{} {
	values := make([]interface{}, maxLevel)
	for i := range values {
		values[i] = convertNumber(float64(baseDur.BaseValue) + float64(i)*baseDur.Increment)
	}
	return values[0:max:max]
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
