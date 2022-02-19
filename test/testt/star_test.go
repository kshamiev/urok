package testt

import (
	"fmt"
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
