package sample

// единоразовое выделение памяти для среза
func sliceGCValid() {
	s := make([]int, 10000)
	for i := 0; i < len(s); i++ {
		s[i] = i
		// какая-то работа
	}
}

// многократное выделение памяти для среза
func sliceGCInvalid() {
	var s []int
	for i := 0; i < 10000; i++ {
		s = append(s, i)
		// какая-то работа
	}
}
