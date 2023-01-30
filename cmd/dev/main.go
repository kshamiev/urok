package main

import (
	_ "fmt"

	uuidc "github.com/gofrs/uuid"
	"github.com/google/uuid"
)

type Funtik struct {
	ID    uuidc.UUID
	Name  string
	Price int64
}

type Fantik struct {
	ID    uuid.UUID
	Nane  string
	Price int
}

//

type Calcer[ID uuid.UUID | uuidc.UUID, P int | int64] interface {
	~struct {
		ID    ID
		Nane  string
		Price P
	}
	Calc()
}

type Calc[ID uuid.UUID | uuidc.UUID, P int | int64, T Calcer[ID, P]] struct {
	ID    ID
	Nane  string
	Price P
}

func NewCalc[ID uuid.UUID | uuidc.UUID, P int | int64, T Calcer[ID, P]](obj T) Calc[ID, P, T] {
	obj.

	return Calc[ID, P, T]{

	}
}

func (self *Calc[ID, P, T]) Calc() {
	//
}

//

func main() {

}
