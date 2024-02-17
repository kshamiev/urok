package main

import (
	"fmt"
	"runtime"
	"strings"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())

	in := "qwertyasdfghj"

	fmt.Println(strings.Index(in, ""))

}
