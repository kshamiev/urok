// Одиночка
// Реализация единственного экземпляра объекта
// Его сокрытие и механизм работы с ним
package main

import "fmt"

func main() {
	settings := GetInstance()
	settings.Host = "192.168.100.1"
	settings.Port = 33
	settings1 := GetInstance()
	fmt.Println(settings1)
}

// Settings is simple struct
type Settings struct {
	Port int
	Host string
}

var instance *Settings

// GetInstance returns a single instance of the settings
func GetInstance() *Settings {
	if instance == nil {
		instance = &Settings{} // <--- NOT THREAD SAFE
	}
	return instance
}
