package main

import (
	"fmt"
	"sort"
)

func main() {
	var A []int
	var res int

	A = []int{1, 3, 6, 4, 1, 2}
	res = Solution(A)
	fmt.Println(res)

	A = []int{1, 2, 3}
	res = Solution(A)
	fmt.Println(res)

	A = []int{-1, -3}
	res = Solution(A)
	fmt.Println(res)

	A = make([]int, 0, 100000)
	for i := 100000; i > 0; i-- {
		A = append(A, i)
	}
	res = Solution(A)
	fmt.Println(res)

	A = make([]int, 2000001)
	j := -1000000
	for i := 0; i < 2000001; i++ {
		A[i] = j
		j++
	}
	res = Solution(A)
	fmt.Println(res)
}

func Solution(A []int) int {
	res := 1
	sort.Sort(numberSort(A))
	for i := range A {
		if A[i] < 1 {
			continue
		}
		if A[i] == res {
			res++
			continue
		}
		if A[i] > res {
			return res
		}
	}
	return res
}

// ////

type numberSort []int

// Len is the number of elements in the collection.
func (v numberSort) Len() int {
	return len(v)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (v numberSort) Less(i, j int) bool {
	if v[i] < v[j] {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (v numberSort) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
