package main

import "fmt"

func main() {
	aa := []int{
		2, 4, 6,
	}
	fmt.Println(Index(aa, 4))
	fmt.Println(Index(aa, 5))

	aaa := []int{
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	}
	Reverse(aaa)
	fmt.Println(aaa)
	ReverseComparable(aaa)
	fmt.Println(aaa)
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func Reverse[T int | int32 | int64](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func ReverseComparable[T comparable](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}
