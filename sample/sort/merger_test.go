//  go test -v -bench=. merger.go merger_test.go utilites.go
package sort

import "testing"

func BenchmarkSortMerge(b *testing.B) {
	b.Run("1 000", func(b *testing.B) {
		b.ReportAllocs()
		items := make([]int, 1000)
		for i := range items {
			items[i] = int(GenInt(1000))
		}
		b.ResetTimer()
		SortMerge(items)
	})
	b.Run("10 000", func(b *testing.B) {
		b.ReportAllocs()
		items := make([]int, 10000)
		for i := range items {
			items[i] = int(GenInt(10000))
		}
		b.ResetTimer()
		SortMerge(items)
	})
	b.Run("100 000", func(b *testing.B) {
		b.ReportAllocs()
		items := make([]int, 100000)
		for i := range items {
			items[i] = int(GenInt(100000))
		}
		b.ResetTimer()
		SortMerge(items)
	})
}
