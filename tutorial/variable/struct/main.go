// Примеры работы со структурами
//
// Структура хранит данные, но не поведение.
// Интерфейс хранит поведение, но не данные.
package main

import "fmt"

func main() {
	obj := NewItemLn()
	obj.CalcLn1()                                 // &{{1000 Вася Person 222} true 56.45 20}
	fmt.Println(obj, obj.Count, obj.Person.Count) // &{{1000 Popcorn 111} true 56.45 567}
}

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

// конструктор типа
func NewItemLn() *Item {
	return &Item{
		Person: Person{
			ID:    1000,
			Name:  "Popcorn",
			Count: 111,
		},
		Flag:  true,
		Price: 56.45,
		Count: 567,
	}
}

// метод типа (работаем по значению)
// это значит что он просто не меняет исходную переменную - приемник (область видимости тут не причем)
func (o Item) CalcLn1() {
	o.Name = "Вася"
	o.Count = 20
	o.Person.Name = "Вася Person"
	o.Person.Count = 222
}

// метод типа (работаем по ссылке)
// а здесь исходная переменная - приемник меняется
func (o *Item) CalcLn2() {
	o.Name = "Вася"
	o.Count = 20
	o.Person.Name = "Вася Person"
	o.Person.Count = 222
}
