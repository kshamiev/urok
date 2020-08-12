package main

import (
	"fmt"
)

// тип реализующий базовую работу с БД (создание, обновление, загрузка одной модели - boiler, gorm, etc...)
type TestNativeDB struct {
	ID    uint64
	Name  string
	Price float64
}

func (obj *TestNativeDB) Load() {
	obj.ID = 456
	obj.Name = "popcorn"
	obj.Price = 85.36
}

// тип реализующий взаимодействие по GRPC и какие-то общие методы (валидации полей п осценарию к примеру)
type TestBase struct {
	*TestNativeDB
}

// метод сопоставление с типом для работы по GRPC
func (obj *TestBase) Convert() {
	fmt.Println(obj.ID)
	fmt.Println(obj.Name)
	fmt.Println(obj.Price)
}

// конечная бизнес модель
type Test struct {
	*TestBase
	env string
}

//
func main() {
	obj := &Test{TestBase: &TestBase{TestNativeDB: &TestNativeDB{}}}
	obj.Load()
	obj.Convert()
}
