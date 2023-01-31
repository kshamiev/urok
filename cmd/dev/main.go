package main

import (
	"fmt"
	_ "fmt"
)

type Ping struct {
	Name string
}

func (o Ping) Hash() uintptr {
	return uintptr(0)
}

type B int

func (o B) Hash() uintptr {
	return uintptr(0)
}

func main() {

	obj := Ping{}
	test(obj)

	obj1 := B(20)

	test1(obj1)
}

func test[T ComparableHasher](n T) {
	fmt.Println(n)
}

func test1[T ComparableHasher1](n T) {
	fmt.Println(n)
}

// ComparableHasher is a type constraint that matches all
// comparable types with a Hash method.
type ComparableHasher interface {
	comparable
	Hash() uintptr
}
type ComparableHasher1 interface {
	~int
	Hash() uintptr
}

func DoubleDefined[S ~[]E, E int](s S) E {

	var r := make(S, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}
