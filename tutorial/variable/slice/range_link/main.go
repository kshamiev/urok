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

// что отобразится после вызова?
func testData1() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
}

// что отобразится после вызова?
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

// что отобразится после вызова?
func testData3() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		go func() {
			fmt.Printf("%d ", a[i])
		}()
	}
}
