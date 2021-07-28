package main

import "fmt"

func main() {
	newSlice()
	copySlice()
}

func newSlice() {
	index := 3
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	listNew := append(list[0:index], list[index+1:]...)
	fmt.Println(listNew)
}

func copySlice() {
	SlTest2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SlTest3 := make([]int, len(SlTest2), cap(SlTest2))
	copy(SlTest3, SlTest2)
	fmt.Println(SlTest3)
}
