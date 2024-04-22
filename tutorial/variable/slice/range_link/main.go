package main

import (
	"fmt"
	"time"
)

func main() {
	testData1()
	time.Sleep(5 * time.Second)
	fmt.Println()

	testData2()
	time.Sleep(5 * time.Second)
	fmt.Println()

	testData3()
	time.Sleep(5 * time.Second)
}

func testData1() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
}

// v 1.21
// three
// three
// three

// v 1.22
// three
// two
// one

func testData2() {
	a := []int{1, 2, 3, 4}
	result := make([]*int, len(a))
	for i, v := range a {
		result[i] = &v
	}
	for _, u := range result {
		fmt.Printf("%d ", *u)
	}
	*result[0] = *result[0] * 2
	fmt.Println()
	for _, u := range result {
		fmt.Printf("%d ", *u)
	}
	fmt.Println()
}

// v 1.21
// 4 4 4 4
// 8 8 8 8

// v 1.22
// 1 2 3 4
// 2 2 3 4

func testData3() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		go func() {
			fmt.Printf("%d ", a[i])
		}()
	}
}

// v 1.21
// 7 8 8 8 8 8 8 8

// v 1.22
// 1 4 2 3 5 8 7 6
