// Компоновщик
// Эффект матрешки. Работает через интерфейс
// Вкладывает объекты реализующие интерфейс один в другой
// Которые в дальнейшем вызывают последующего вложенного по цепочке через родственный метод
// Не замыкающая рекурсия на дочерний объект
package main

import "fmt"

func main() {
	image := Image{}
	image.Add(Сircle{})
	image.Add(Square{})
	picture := Image{}
	picture.Add(image)
	picture.Add(Image{})
	picture.Draw()
}

// Graphic is Component
type Graphic interface {
	Draw()
}

// Сircle is Leaf
type Сircle struct{}

// Draw is Operation
func (c Сircle) Draw() {
	fmt.Println("Draw circle")
}

// Square is Leaf
type Square struct{}

// Draw is Operation
func (s Square) Draw() {
	fmt.Println("Draw square")
}

// Image is Composite
type Image struct {
	graphics []Graphic
}

// Add Adds a Leaf to the Composite.
func (i *Image) Add(graphic Graphic) {
	i.graphics = append(i.graphics, graphic)
}

// Draw is Operation
func (i Image) Draw() {
	fmt.Println("Draw image")
	for _, g := range i.graphics {
		g.Draw()
	}
}
