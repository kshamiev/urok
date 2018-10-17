// Примеры работы со структурами
//
// Структура хранит данные, но неповедение.
// Интерфейс хранит поведение, но не данные.
package structandface

import (
	"fmt"
)

type Person struct {
	ID    uint64
	Name  string
	Count int64
}

type Item struct {
	Person
	Flag  bool
	Price float64
	Count int64
}

// конструктор типа и возможные варианты его инициализации
func NewItemLn() *Item {

	// return new(Item)

	obj := &Item{
		Flag:  true,
		Price: 56.45,
		Count: 567,
	}
	obj.ID = 1000
	obj.Name = "Popcorn"
	obj.Person.Count = 111
	return obj

	// встраиваемые типы так инициализировать не получиться
	return &Item{
		Flag:  true,
		Price: 56.45,
		Count: 567,
	}

	// return &Item{true, 56.45, 3567} // встраиваемые типы так инициализировать не получиться
}

// метод типа (работаем по ссылке)
func (o *Item) CalcLn() {
	o.Count = 20
	o.Name = "Вася"
	o.Person.Name = "Вася Person"
	o.Person.Count = 222
}

// Пример
func SampleStruct() {

	//	var obj = NewItemZn()
	//	fmt.Println("по значению: ", obj)
	var obj = NewItemLn()
	fmt.Println("по ссылке: ", obj)

	//	obj.CalcZn()
	//	fmt.Println(obj)
	obj.CalcLn()
	fmt.Println(obj)
	fmt.Println(obj.Name)
	fmt.Println(obj.Count)
	fmt.Println(obj.Person.Name)
	fmt.Println(obj.Person.Count)

	//	sampleStructLink(obj)
	//	fmt.Println(obj)
}

func sampleStructLink(obj *Item) {
	obj.Count = 100
}

func NewItemZn() Item {

	obj := Item{
		Flag:  true,
		Price: 56.45,
		Count: 567,
	}

	obj.ID = 1000
	obj.Name = "Popcorn"
	obj.Person.Count = 111

	return obj

	// встраиваемые типы так инициализировать не получиться
	return Item{
		Flag:  true,
		Price: 56.45,
		Count: 567,
	}

	//	return Item{}
}

func (o Item) CalcZn() {
	o.Count = 10
}
