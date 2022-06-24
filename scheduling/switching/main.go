package main

import (
	"strings"
)

func main() {

}

// Подсчёт целых чисел
// разделение работы по нескольким горутинам
// плюс увеличение количества потоков даст существенный прирост производительности
func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}

// Сортировка пузырьком
// разделение выполнение задачи на горутины не даст прироста пр. из-за специфики алгоритма и задачи
func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, currentPass int) bool {
	var idx int
	idxNext := idx + 1
	n := len(numbers)
	var swap bool

	for idxNext < (n - currentPass) {
		a := numbers[idx]
		b := numbers[idxNext]
		if a > b {
			numbers[idx] = b
			numbers[idxNext] = a
			swap = true
		}
		idx++
		idxNext = idx + 1
	}
	return swap
}

// Чтение файлов и поиск по тексту
// разделение работы по нескольким горутинам даст существенный прирост производительности
// количество потоков не влияет на производительность так как происходит естественное переключение между горутинами
func find(topic string, docs []string) int {
	var found int
	for _, doc := range docs {
		items, err := read(doc)
		if err != nil {
			continue
		}
		for _, item := range items {
			if strings.Contains(item, topic) {
				found++
			}
		}
	}
	return found
}

func read(doc string) ([]string, error) {
	// реализация чтение документа из FS
	return []string{}, nil
}
