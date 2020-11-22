// Прототип
// Создание (порождение) самого себя со всем своим состоянием
// Или другими словами клонирование самого себя
package main

import "fmt"

func main() {
	square := Square{4}
	maker := ShapeMaker{square}

	square1 := maker.MakeShape()
	square2 := maker.MakeShape()

	fmt.Println(square1)
	fmt.Println(square2)
	fmt.Println(square.Clone())
}

// Shape is Prototype
type Shape interface {
	Clone() Shape
}

// Square is ConcretePrototype
type Square struct {
	LineCount int
}

// Clone creates a copy of the square
func (s Square) Clone() Shape {
	return Square{s.LineCount}
}

// ////

// ShapeMaker contains a Shape
type ShapeMaker struct {
	Shape Shape
}

// MakeShape creates a copy of the Shape
func (sm ShapeMaker) MakeShape() Shape {
	return sm.Shape.Clone()
}
