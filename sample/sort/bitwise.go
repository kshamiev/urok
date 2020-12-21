// Алгоритм поразрядной сортировки
//
package sort

import (
	"math"
)

// Time Complexity O(nk)
// Space Complexity O(n+k)
func SortBitwise(arr []int, cBase int) []int {
	maxVal := 0
	for _, value := range arr {
		if value > maxVal {
			maxVal = value
		}
	}

	i := 0
	for math.Pow(float64(cBase), float64(i)) <= float64(maxVal) {
		arr = bucketsToList(listToBuckets(arr, cBase, i))
		i++
	}
	return arr
}

func listToBuckets(items []int, cBase int, i int) [][]int {
	var buckets = make([][]int, cBase)

	var pBase = int(math.Pow(
		float64(cBase), float64(i)))
	for _, x := range items {
		// Isolate the base-digit from the number
		var digit = (x / pBase) % cBase
		// Drop the number into the correct bucket
		buckets[digit] = append(buckets[digit], x)
	}

	return buckets
}

func bucketsToList(buckets [][]int) []int {
	result := []int{}

	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}
