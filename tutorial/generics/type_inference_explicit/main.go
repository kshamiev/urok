package main

import (
	"fmt"
)

// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 |
		~int64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

// Sum returns the sum of the provided arguments.
func Sum[T Numeric](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// id is a type alias for an int64
type id int64

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum([]id{1, 2, 3}...))

	// Generic types can be specified explicitly by invoking a function
	// with the bracket notation and the list of types to use. Because
	// the Sum function only has a single, generic type -- "T" -- the
	// call "Sum[float64]" means that "T" will be replaced by "float64"
	// when compiling the code. Since the values "1" and "2" can both
	// be treated as "float64," the code is valid.
	fmt.Println(Sum[float64](1, 2, 3.0))

	// Вывод типа или явное указание типа:
	// Указание типа перед вызовом функции необходимо в случаях когда на вход передаются значения разных типов
	// (разрешённых в составном ограничительном типе)
	// чтобы компилятор однозначно определил нужный тип при сборке программы
}
