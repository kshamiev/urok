package sort

import (
	"strconv"
	"testing"
)

// GOGC=off go test ./sample/sort/. -run=^# -bench=SortMerge -benchtime=10x -count 3
func BenchmarkSortMerge(b *testing.B) {
	for _, number := range []int{1000, 10000, 100000} {
		b.Run(strconv.Itoa(number), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.ReportAllocs()
				b.SetBytes(2)
				items := make([]int, number)
				for i := range items {
					items[i] = int(GenInt(int64(number)))
				}
				b.ResetTimer()
				SortMerge(items)
			}
		})
	}
}
