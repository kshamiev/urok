// Алгоритм сортировки пузырьком
//
package sort

// Time Complexity O(n^2)
// Space Complexity O(1)
func SortBubble(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[i] {
				var tmp = arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
}
