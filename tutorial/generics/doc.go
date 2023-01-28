package generics

func Test1[T int | int16 | int32 | int64](a, b T) T {
	return a + b
}

func Test2[T comparable](a, b T) bool {
	return a == b
}

type Test[T int | int64 | float64] struct {
	Price T
}

// T
// generic - (Общий) Это символ или мета-тип, представляющий один или несколько конкретных типов или интерфейсов

// int | int16 | int32 | int64
// это ограничение, указывающее, какие конкретные типы можно использовать.

// ~
// тильда означает разрешение ограничения всех производных от указанного за ним типа

// any
// синоним interface{}

// comparable
// это интерфейс-ограничение сопоставимых базовых типов и структур

// Составные ограничение типа (вынос ограничение в отдельный тип)

type NumberConstraint interface {
	int64 | float64 | ~int
}
