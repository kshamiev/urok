package main

import (
	"fmt"
	_ "fmt"

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
}

func G[T Funtik | Rrrr, K ~int | ~int64](o Avg[T, K]) {
	o.Name = "ffff"
	fmt.Printf("%T\n", o)
	fmt.Printf("%T\n", o.Obj)
}
