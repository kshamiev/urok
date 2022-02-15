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

func SampleFace() {
	sampleFly(&Bird{"Чайка"})
	fmt.Println()
	sampleFly(&Aircraft{"Mig-31"})
}

type flying interface {
	Fly()
}

// пример передачи значение типа через интерфейс который он реализует
// и использование утверждение типа
func sampleFly(obj flying) {

	// это абстрактный механизм проверки на реальный nil (но работает не быстро)
	// быстрый (по производительности) способ проверки на nil это реализовывать методы у интерфейса и объектов которые делают эту проверку
	if !reflect.ValueOf(obj).IsNil() {
		fmt.Println("type check is nukll")
		obj.Fly()
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

// ////

// для понимания отображения в панели навигации
type FlyingPub interface {
	Test()
}

// для понимания отображения в панели навигации
type flyingPubTest struct {
	Name string
}
