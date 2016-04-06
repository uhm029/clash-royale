package lib

var (
	V1000R = []int{1000, 1100, 1210, 1330, 1460, 1600, 1760, 1930, 2120, 2330, 2560, 2810}
)

func generateHp(value interface{}) []int {
	baseValue := value.(int)
	values := make([]int, len(V1000R))
	for i, v := range V1000R {
		values[i] = v * baseValue / 1000
	}
	return values
}

func generateDam(value interface{}) []int {
	return generateHp(value)
}
