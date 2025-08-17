package five

import (
	"math"
)

// площадь
// S - площадь a - длина b - ширина

func square(a, b int) int {
	S := a * b
	return S
}

// периметр
// P - периметр a - длина b - ширина

func perimeter(a, b int) int {
	P := (a + b) * 2
	return P
}

// движение
// S - путь v - скорость t - время

func motion(v, t int) int {
	S := v * t
	return S
}

// Объём прямоугольного параллелепипеда.
// V - объём a, b, c стороны (ДШВ)

func parallelepiped(a, b, c int) int {
	V := a * b * c
	return V
}

// Объём наклонного параллелепипеда.
// Объём любого параллелепипеда.

func parallelepipedAll1(a, b, h int) int {
	V := a * b * h
	return V
}

// w угол наклона C к основанию A и B

func parallelepipedAll2(a, b, c int, w float64) int {
	V := a * b * c * int(math.Sin(w))
	return V
}
