package main

import (
	"fmt"
	"time"
)

func main() {
	testGoroutines3()
	testData2()
}

// что отобразится после вызова?
func testGoroutines3() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
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
}
