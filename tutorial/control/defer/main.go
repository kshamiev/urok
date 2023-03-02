package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := Work()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
}

func Work() error {
	fp, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o666)
	if err != nil {
		return err
	}
	defer fp.Close()

	fp.WriteString("Фиалка Губоцветная\n")
	// It`s Work
	return nil
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func cc() int {
	i := 0
	defer func() { i++ }()
	return 1
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
