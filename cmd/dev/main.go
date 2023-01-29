package main

import (
	"fmt"
	_ "fmt"
	"strings"

	"github.com/google/uuid"
)

type Funtik struct {
	ID    uuid.UUID
	Name  string
	Price int64
}

type Rrrr struct {
	ID    uuid.UUID
	Nane  string
	Price int
}

type Avg[T Funtik | Rrrr, K ~int | ~int64] struct {
	Obj   T
	Name  string
	Price K
}

func main() {
	G(Avg[Funtik, int64]{Obj: Funtik{}})
	G(Avg[Rrrr, int]{})
	Ffff([]int{1, 2, 3, 4, 5})
}

func G[T Funtik | Rrrr, K ~int | ~int64](o Avg[T, K]) {
	o.Name = "ffff"
	fmt.Printf("%T\n", o)
	fmt.Printf("%T\n", o.Obj)
}

type Vector[T any] []T

func (v *Vector[T]) Push(x T) { *v = append(*v, x) }

var v Vector[int]

func Ffff[T any](l Vector[T]) {
	l.Push(l[0])
	for i := range l {
		fmt.Printf("%T %v\n", l[i], l[i])
	}
	fmt.Printf("%T\n", v)
}

type Stringer interface {
	String() string
}

type StringableVector[T Stringer] []T

func (s StringableVector[T]) String() string {
	var sb strings.Builder
	for i, v := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		// It's OK to call v.String here because v is of type T
		// and T's constraint is Stringer.
		sb.WriteString(v.String())
	}
	return sb.String()
}
