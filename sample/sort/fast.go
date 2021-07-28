// Алгоритм быстрой сортировки
//
// Находит значение элемента находящегося по середине
// Далее ведет поиск максимального значения с лева и минимального значения справа относительно найденного
// Найдя меняет их местами
// Запоминает их индексы
// Далее идет рекурсивно на следующую итерацию сортировки с учетом этих индексов
//
// И сортировка продолжается
// каждую итерацию сортировка делится на левую и и правую часть от полученных индексов (с их пограничным пересечением)
//
package sort

// Time Complexity from O(n*log(n)) to O(n^2)
// Space Complexity O(log(n))
func SortQuick(items []int, fst int, lst int) {
	if fst >= lst {
		return
	}
	i := fst
	j := lst
	x := items[(fst+lst)/2]
	for i < j {
		for items[i] < x {
			i++
		}
		for items[j] > x {
			j--
		}
		if i <= j {
			var tmp = items[i]
			items[i] = items[j]
			items[j] = tmp
			i++
			j--
		}
	}
	SortQuick(items, fst, j)
	SortQuick(items, i, lst)
}
