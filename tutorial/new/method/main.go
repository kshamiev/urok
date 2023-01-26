package main

import (
	"fmt"

	"github.com/kshamiev/urok/tutorial/new/typ"
)

func main() {
	obj := &General[typ.Good]{
		Description: "Описание",
	}
	_ = obj.Test(&typ.User{})
	_ = obj.Update(typ.Good{})
	_ = obj.UpdatePoint(&typ.Good{})
	fmt.Println()

	obj1 := &General[typ.Invoice]{
		Description: "Описание",
	}
	_ = obj1.Test(&typ.User{})
	_ = obj1.Update(typ.Invoice{})
	_ = obj1.UpdatePoint(&typ.Invoice{})
}

type General[T typ.Good | typ.Invoice] struct {
	Description string
}

func (obj *General[T]) Test(item *typ.User) error {
	fmt.Printf("%T\n", item)
	return nil
}

func (obj *General[T]) Update(item T) error {
	fmt.Printf("%T\n", item)
	return nil
}

func (obj *General[T]) UpdatePoint(item *T) error {
	fmt.Printf("%T\n", item)
	return nil
}
