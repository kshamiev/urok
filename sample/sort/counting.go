// Алгоритм сортировки подсчетом
//
package sort

import (
	"math"
)

// Time Complexity O(n+k)
// Space Complexity O(k)
func SortCounting(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)

	var min = math.MaxInt32
	var max = math.MinInt32
	for _, x := range arr {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	var counts = make([]int, max-min+1)

	for _, x := range arr {
		counts[x-min]++
	}

	var total = 0
	for i := min; i <= max; i++ {
		var oldCount = counts[i-min]
		counts[i-min] = total
		total += oldCount
	}

	for _, x := range arr {
		items[counts[x-min]] = x
		counts[x-min]++
	}
	return items
}
