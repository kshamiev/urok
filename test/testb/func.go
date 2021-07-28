package testb

// ////

// O(N^2)
func AlgorithmComplexity1(n int, data map[int]map[int]int) map[int]int {
	var res = make(map[int]int)
	for i := 0; i < n; i++ {
		res[i] = data[i][0]
		for j := 0; j < n; j++ {
			if res[i] < data[i][j] {
				res[i] = data[i][j]
			}
		}
	}
	return res
}

func GetAlgorithmComplexity1(n int) map[int]map[int]int {
	data := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		data[i] = make(map[int]int)
		for j := 0; j < n; j++ {
			data[i][j] = int(GenInt(1000000))
		}
	}
	return data
}
