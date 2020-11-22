// Фабричный метод
// Не правильная реализация
// Он должен создавать сам себя
// Этот паттерн среднее между фабрикой и прототипом
package main

import "fmt"

func main() {
	booker := BookerCreator{}.CreateEmployee()
	booker.Test()
	// printed: Booker

	manager := ManagerCreator{}.CreateEmployee()
	manager.Test()
	// printed: Manager
}

// Employee is Product
type Employee interface {
	Test()
}

// Booker is ConcreteProduct
type Booker struct{}

// Test is Employee method
func (e Booker) Test() {
	fmt.Println("Booker")
}

// Manager is ConcreteProduct
type Manager struct{}

// Test is Manager method
func (e Manager) Test() {
	fmt.Println("Manager")
}

// BookerCreator is ConcreteCreator
type BookerCreator struct{}

// CreateEmployee creates an Booker
func (c BookerCreator) CreateEmployee() Employee {
	return Booker{}
}

// ManagerCreator is ConcreteCreator
type ManagerCreator struct{}

// CreateEmployee creates an Manager
func (c ManagerCreator) CreateEmployee() Employee {
	return Manager{}
}
