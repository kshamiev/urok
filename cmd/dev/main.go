package main

// https://stackoverflow.com/questions/71418671/restart-or-shutdown-golang-apps-programmatically

import (
	"time"
)

func main() {
	gg := []string{"1", "2", "3", "4", "5"}
	for j := 0; j < 1000; j++ {
		go func() {
			for {
				for i := range gg {
					_ = gg[i]
				}
			}
		}()
	}
	time.Sleep(time.Minute)
}
