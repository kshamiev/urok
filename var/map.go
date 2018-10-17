package main

import (
	"fmt"
)

var Hash0 map[string]string         //
var Hash1 = make(map[string]string) // рекомендуется
var Hash2 = map[string]string{
	"key1": "value1",
	"key2": "value2",
}                                   // рекомендуется

func main() {
	//	Hash0["qwerty"] = "yuiop" // здесь будет ошибка, так как хеш просто обьявлен и имеет значение nil
	Hash1["qwerty1"] = "yuiop1"
	fmt.Println("Хеши: ", Hash0, Hash1, Hash2, "\n")

	sampleMap(Hash1)
	fmt.Println(Hash1, "\n")

	if _, ok := Hash1["qwerty1"]; ok == true {
		fmt.Println("Хеш index qwerty1 OK", "\n")
	} else {
		fmt.Println("Хеш index qwerty1 NOT", "\n")
	}

	delete(Hash1, "qwerty1")

	if _, ok := Hash1["qwerty1"]; ok == true {
		fmt.Println("Хеш index qwerty1 OK", "\n")
	} else {
		fmt.Println("Хеш index qwerty1 NOT", "\n")
	}
}

// Хеши всегда передаются по ссылке
func sampleMap(hash map[string]string) {

	hash["element"] = "now element"

}
