// Примеры работы с интерфейсами
// Утиная типизация это когда тип имеет методы удовлетворяющие всем методам описанным в конкретном интерфейсе
// Таким образом без како-го либо явного указания данный тип является также типом интерфейсом
//
// Структура хранит данные, но не поведение.
// Интерфейс хранит поведение, но не данные.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var o *Bird
	sampleFly(o)
	fmt.Println()
	sampleFly(&Bird{"Чайка"})
	fmt.Println()
	sampleFly(&Aircraft{"Mig-31"})
}

// пример передачи значение типа через интерфейс который он реализует
// и использование утверждение типа
func sampleFly(obj flying) {

	if obj == nil {
		fmt.Println("type check is nil")
		return
	}

	// это абстрактный механизм проверки на реальный nil (но работает не быстро)
	// быстрый (по производительности) способ проверки на nil это реализовывать методы у интерфейса и объектов которые делают эту проверку
	if reflect.ValueOf(obj).IsNil() {
		fmt.Println("type check is reflect.nil")
		return
	}

	obj.Fly()
	if o, ok := obj.(*Bird); ok { // утверждение типа
		fmt.Println("type assertion through IF", o.Name)
	}
	if o, ok := obj.(*Aircraft); ok { // утверждение типа
		fmt.Println("type assertion through IF", o.Name)
	}

	switch o := obj.(type) { // утверждение типа
	case *Bird:
		fmt.Println("type assertion through SWITCH", o.Name)
	case *Aircraft:
		fmt.Println("type assertion through SWITCH", o.Name)
	}
}

// проверка что тип реализует указанный интерфейс
var _ flying = &Bird{}
var _ flying = &Aircraft{}
var _ flying = (*Aircraft)(nil)

// var _ flying = &AircraftFail{}
// var _ flying = (*AircraftFail)(nil)

type flying interface {
	Fly()
}

type Bird struct {
	Name string
}

func (o *Bird) Fly() {
	fmt.Println("fly is *Bird: ", o.Name)
}

type Aircraft struct {
	Name string
}

func (o *Aircraft) Fly() {
	fmt.Println("fly is *Aircraft: ", o.Name)
}
