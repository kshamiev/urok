// Декоратор (похож на интерпретатор)
//
// Тип обертка над другим типом с альтернативной реализацией его функционала но с соблюдением  сигнатуры
package main

import "fmt"

func main() {
	square := Square{}
	square.ShowInfo()
	// printed: square
	fmt.Println()

	colorShape := ColorShape{ShapeDecorator{square}, "red"}
	colorShape.ShowInfo()
	// printed: red square
	fmt.Println()

	shadowShape := ShadowShape{ShapeDecorator{colorShape}}
	shadowShape.ShowInfo()
	// printed: red square with shadow
}

// Shape is Component
type Shape interface {
	ShowInfo()
}

// ShapeDecorator is Decorator
type ShapeDecorator struct {
	Shape Shape
}

// ShowInfo is Operation()
func (sd ShapeDecorator) ShowInfo() {
	sd.Shape.ShowInfo()
}

// ////

// Square is ConcreteComponent
type Square struct{}

// ShowInfo is Operation()
func (s Square) ShowInfo() {
	fmt.Print("square")
}

// ColorShape is ConcreteDecorator
type ColorShape struct {
	ShapeDecorator
	color string
}

// ShowInfo is Operation()
func (cs ColorShape) ShowInfo() {
	fmt.Print(cs.color + " ")
	cs.Shape.ShowInfo()
}

// ShadowShape is ConcreteDecorator
type ShadowShape struct {
	ShapeDecorator
}

// ShowInfo is Operation()
func (ss ShadowShape) ShowInfo() {
	ss.Shape.ShowInfo()
	fmt.Print(" with shadow")
}
