package main

import (
	"fmt"

	"github.com/kshamiev/urok/tutorial/new/typ"
)

func main() {
	_ = Update(typ.Good{})
	_ = Update(typ.Invoice{})
	fmt.Println()

	_ = UpdatePoint(&typ.Good{})
	_ = UpdatePoint(&typ.Invoice{})
	fmt.Println()

	a := MyInt(2)
	b := MyInt(3)
	fmt.Println(add1(a, b))
	fmt.Println(add2(a, b))

	aa := []int{
		2, 4, 6,
	}
	bb := 4
	fmt.Println(Index(aa, bb))

	aaa := []int{
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	}
	Reverse(aaa)
	fmt.Println(aaa)

}

func Update[T typ.Good | typ.Invoice](item T) error {
	fmt.Printf("%T\n", item)
	return nil
}

func UpdatePoint[T typ.Good | typ.Invoice](item *T) error {
	fmt.Printf("%T\n", item)
	return nil
}

// 1
type MyInt uint64

func add1[T int | ~uint64](a T, b T) T {
	return a + b
}

type additive interface {
	int | ~uint64
}

func add2[T additive](a T, b T) T {
	return a + b
}

// 1

// 2
// comparable
// ==
// !=
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// 2

func Reverse[T int | int32 | int64](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// ////

type Stringer interface {
	String() string
}

func Tos[T Stringer](s []T) []string {
	var ret []string
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

func TosTos(s []Stringer) []string {
	var ret []string
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

func Scale[S ~[]E, E int](s S, sc E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * sc
	}

	return r
}
