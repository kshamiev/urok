package sort

import (
	"strconv"
	"testing"
)

/*
-run=^#
-run none
отключение запуска обычных или модульных тестов

-bench=SortBitwise
фильтр, какой тест будет запущен

-benchtime=10x
количество итераций внутри теста (b.N)
-benchtime=10s
время в секундах на тест

-benchmem alias b.ReportAllocs()
показывать аллокации на куче и время выполнения одной итерации внутри теста (b.N)

-count 3
количество запусков теста

-cpu 8
количество используемых процессоров

GOGC=off
Отключение GC
*/

// GOGC=off go test ./sample/sort/. -run=^# -bench=SortBitwise -benchtime=10x -count 3 -cpu 8
func BenchmarkSortBitwise(b *testing.B) {
	for _, number := range []int{1000, 10000, 100000} {
		b.Run(strconv.Itoa(number), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.ReportAllocs()
				b.SetBytes(2)
				b.ResetTimer()
				b.StopTimer()
				items := make([]int, number)
				for i := range items {
					items[i] = int(GenInt(int64(number)))
				}
				b.StartTimer()
				SortBitwise(items, 10)
			}
		})
	}
}

/*
BenchmarkSortBitwise/1000-8	    10    9971 ns/op    0.20 MB/s    11606 B/op    25 allocs/op

BenchmarkSortBitwise/1000-8
имя теста - количество используемых потоков (логических П)
10
количество итераций тестирования функции

9971 ns/op
время затраченное в среднем на каждую итерацию в наносекундах.

0.20 MB/s
пропускная способность за одну итерацию в мегабайтах
или количество обработанной памяти за одну итерацию
(это то сколько данных было обработано)

11606 B/op
память выделенная на одну итерацию в байтах
(это сколько памяти было выделено)

25 allocs/op
количество выделений памяти на куче не на стеке
(это количесвто обращений к куче для выделения памяти)

*/
