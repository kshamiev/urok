package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kshamiev/urok/temp/jsn"
)

func main() {

	o := &Test{
		ID:    34,
		Fio:   `Pupkin,:K`,
		Price: 45.78,
		Name:  `"sdfsdfsd"`,
		Funtik: TestChild{
			ID:   546,
			Name: "hgjhjhjklfgh",
		},
		Fantik: []TestChild{
			{
				ID:   546,
				Name: "hgjhjhjklfgh",
			},
			{
				ID:   546,
				Name: "hgjhjhjklfgh",
			},
		},
	}

	data, err := json.Marshal(o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	if err := json.Unmarshal(data, o); err != nil {
		log.Fatal(err)
	}
}

type Test struct {
	ID     int
	Fio    string
	Price  float64
	Name   string
	Fantik []TestChild
	Funtik TestChild
}

// func (o Test) MarshalJSON() ([]byte, error) {
// 	fmt.Println("MarshalJSON !")
// 	return []byte("{}"), nil
// }
//
func (o *Test) UnmarshalJSON(data []byte) error {
	sc := jsn.NewJson(data)
	name, value := sc.Get()
	fmt.Println(name, value)
	name, value = sc.Get()
	fmt.Println(name, value)
	name, value = sc.Get()
	fmt.Println(name, value)
	name, value = sc.Get()
	fmt.Println(name, value)
	name, value = sc.Get()
	fmt.Println(name, value)

	return nil
}

// // //

type TestChild struct {
	ID   int
	Name string
}

// ////

// func (o TestChild) MarshalJSON() ([]byte, error) {
// 	fmt.Println("MarshalJSON !!!")
// 	return []byte("{}"), nil
// }
//
// func (o *TestChild) UnmarshalJSON([]byte) error {
// 	fmt.Println("UnmarshalJSON !!!")
// 	return nil
// }

// Чтение из БД
func (o *Test) Scan(value interface{}) error {
	return nil
}

// Запись в БД
func (o Test) Value() (driver.Value, error) {
	return nil, nil
}

// for i := range data {
// fmt.Println(data[i])
// if data[i] == 123 {
// fmt.Println("{")
// }
// if data[i] == 125 {
// fmt.Println("}")
// }
// if data[i] == 34 {
// fmt.Println(`"`)
// }
// if data[i] == 58 {
// fmt.Println(":")
// }
// if data[i] == 44 {
// fmt.Println(",")
// }
// if data[i] == 92 {
// fmt.Println(`\`)
// }
// if data[i] == 91 {
// fmt.Println(`[`)
// }
// if data[i] == 93 {
// fmt.Println(`]`)
// }
// }
