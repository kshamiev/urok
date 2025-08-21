package five

import (
	"math"
)

// Площадь
// S - площадь a - длина b - ширина

func square(a, b int) int {
	S := a * b
	return S
}

// Периметр
// P - периметр a - длина b - ширина

func perimeter(a, b int) int {
	P := (a + b) * 2
	return P
}

// Движение
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

func parallelepipedAll2(a, b, c, w float64) float64 {
	V := a * b * c * math.Sin(w)
	return V
}
