// Замыкание
// Это когда область видимости (переменных) конкретной функции простирается за ее пределы.
// Такие функции (зачастую анонимные) обьявляются внутри другой функции.
// Таким образом такая функция получает доступ к перменным окружающей ее функции.

// Замыканием называют функцию, использующую переменные, определенные за ее пределами.
// В нашем случае функция increment и переменная x образуют замыкание.

package function

import (
	"fmt"
)

func Closure() {

	x := 7

	// анонимная функция
	increment := func() int {
		x++
		return x
	}
	d := func() {
		fmt.Println(increment())
	}

	fmt.Println(increment())
	fmt.Println(increment())
	d()

	// особенность данного подхода сохранение значения переменных (i) между вызовами функции
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
