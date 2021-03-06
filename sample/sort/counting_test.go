//  go test -v -bench=. counting.go counting_test.go utilites.go
package sort

import "testing"

func BenchmarkSortCounting(b *testing.B) {
	b.Run("1 000", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(2)
		b.ResetTimer()
		b.StopTimer()
		items := make([]int, 1000)
		for i := range items {
			items[i] = int(GenInt(1000))
		}
		b.StartTimer()
		b.Log(items[:10])
		b.Log(items[990:])
		items = SortCounting(items)
		b.Log(items[:10])
		b.Log(items[990:])
	})
	b.Run("10 000", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(2)
		b.ResetTimer()
		b.StopTimer()
		items := make([]int, 10000)
		for i := range items {
			items[i] = int(GenInt(10000))
		}
		b.StartTimer()
		items = SortCounting(items)
	})
	b.Run("100 000", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(2)
		b.ResetTimer()
		b.StopTimer()
		items := make([]int, 100000)
		for i := range items {
			items[i] = int(GenInt(100000))
		}
		b.StartTimer()
		items = SortCounting(items)
	})
}
