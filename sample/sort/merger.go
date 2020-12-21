// Алгоритм сортировки слиянием
//
package sort

// Time Complexity O(n*log(n)))
// Space Complexity O(n)
func SortMerge(items []int) {
	length := len(items)
	if length == 1 {
		return
	}

	var lLeft = length / 2
	var left = make([]int, lLeft)
	copy(left, items[:lLeft])
	var lRight = length - lLeft
	var right = make([]int, lRight)
	copy(right, items[lLeft:])

	SortMerge(left)
	SortMerge(right)

	merge(left, right, items)
}

func merge(left []int, right []int, result []int) {
	l := 0
	r := 0
	i := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	var length = len(left) - l
	copy(result[i:i+length], left[l:])
	i = i + length
	length = len(right) - r
	copy(result[i:i+length], right[r:])
}
