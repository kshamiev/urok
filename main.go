package main

import (
	"fmt"
	"time"

	uuidn "github.com/gofrs/uuid/v5"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("V4")
	for i := 0; i < 10; i++ {
		fmt.Println(uuid.New().String())
		time.Sleep(time.Millisecond)
	}
	fmt.Println("V7")
	for i := 0; i < 10; i++ {
		u, _ := uuidn.NewV7()
		fmt.Println(u.String())
		time.Sleep(time.Millisecond)
	}
}
