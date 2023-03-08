package testt

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSlice(t *testing.T) {
	data := []int{1, 2, 3, 4}
	data = append(data, 666)
	fmt.Println(cap(data))
	sliceFN(&data)
	fmt.Println(data)
}

func sliceFN(in *[]int) {
	*in = append(*in, 125)
	for i := range *in {
		(*in)[i] = 5
	}
	fmt.Println(in)

	fp, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	fp.WriteString("Фиалка Губоцветная\n")
}

// ////

type A struct {
	Name string
}

func (a *A) Fun() {}

type Good interface {
	Fun()
}

func TestFace(t *testing.T) {
	var a *A
	fmt.Println(a)
	faceFN(a)
}

func faceFN(in Good) {
	fmt.Println(in)
	if in == nil {
		fmt.Println("IS NIL")
	} else {
		fmt.Println("IS NOT NIL")
	}
}

// ////
