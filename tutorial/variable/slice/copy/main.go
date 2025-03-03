package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("origin")
	copySlice()
	fmt.Println("cut")
	cutSlice()
	fmt.Println("paste")
	pasteSlice()
	fmt.Println("paste vector")
	res := Insert([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 1, 2, 3)
	fmt.Println(res)
	res = Insert([]int{1, 2, 3}, 5, 1, 2, 3)
	fmt.Println(res)
	res1 := Insert([]string{"1", "2", "3"}, 5, "1", "2", "3")
	fmt.Println(res1)

	res10 := slices.Insert([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 1, 2, 3)
	fmt.Println(res10)
	res10 = slices.Insert([]int{1, 2, 3}, 3, 1, 2, 3)
	fmt.Println(res10)
	res11 := slices.Insert([]string{"1", "2", "3"}, 2, "1", "2", "3")
	fmt.Println(res11)

	res10 = slices.Insert([]int{1, 2, 3}, 1, 9)
	fmt.Println(res10)
}

func copySlice() {
	SlTest2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SlTest3 := make([]int, len(SlTest2), cap(SlTest2))
	copy(SlTest3, SlTest2)
	fmt.Println(SlTest3)
}

func cutSlice() {
	index := 3
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	listNew := append(list[:index-1], list[index:]...)
	fmt.Println(listNew)
}

func pasteSlice() {
	n := 0
	index := 3
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	listNew := append(list[:index+1], list[index:]...)
	listNew[index] = n
	fmt.Println(listNew)
}

func Insert[T any](a []T, index int, b ...T) []T {
	if n := len(a) + len(b); n <= cap(a) {
		s2 := a[:n]
		copy(s2[index+len(b):], a[index:])
		copy(s2[index:], b)
		return s2
	}
	s2 := make([]T, len(a)+len(b))
	if index > len(a) {
		index = len(a)
	}
	copy(s2, a[:index])
	copy(s2[index:], b)
	copy(s2[index+len(b):], a[index:])
	return s2
}

// a = Insert(a, i, b...)
