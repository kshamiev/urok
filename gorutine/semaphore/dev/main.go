package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var n = 1000
	var m = 10000
	var data = make(map[string][]int64)
	for i := 0; i < n; i++ {
		row := make([]int64, 0, m)
		for j := 0; j < m; j++ {
			row = append(row, rand.Int63())
		}
		data[fmt.Sprintf("%d", rand.Int63())] = row
	}
	result := semaphore(data, rand.Intn(100))
	for k, v := range result {
		fmt.Printf("k:%s v:%d\n", k, v)
	}
}

const maxGR = 4

func semaphore(src map[string][]int64, threads int) map[string]int64 {

	concurentGR := make(chan struct{}, threads)
	res := make(map[string]int64, len(src))

	for str, arr := range src {
		concurentGR <- struct{}{}

		go sumArray(str, arr, res)

		<-concurentGR
	}

	wg.Wait()

	return res
}

var mu sync.Mutex
var wg sync.WaitGroup

func sumArray(str string, arr []int64, res map[string]int64) {
	wg.Add(1)
	defer wg.Done()
	r := int64(0)
	for _, n := range arr {
		r += n
	}
	mu.Lock()
	defer mu.Unlock()
	res[str] = r
}
