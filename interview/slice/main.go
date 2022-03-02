package main

import (
	"fmt"
	"time"
)

func main() {
	testGoroutines3()
	testData2()
	testData4()
}

// что отобразится после вызова?
func testGoroutines3() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(2 * time.Second)
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
func testData4() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		go func() {
			fmt.Printf("%d ", a[i])
		}()
	}
	time.Sleep(time.Second * 5)
}
