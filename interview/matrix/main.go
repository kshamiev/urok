/*
Есть матрица NxN, состоящая из 0 и 1,
и отражающая расположения кораблей на поле для морского боя.

1. Кораблей может быть любое количество
2. Размер кораблей — от 1х1 до 1хN
3. Корабли никак не соприкасаются друг с другом.
4. Корабли располагаются горизонтально и вертикально

Необходимо подсчитать количество кораблей.

Пример:

[
  [1, 1, 0, 0, 1, 0},
  [0, 0, 0, 0, 1, 0],
  [1, 0, 1, 0, 1, 0],
  [0, 0, 0, 0, 0, 0],
  [1, 0, 1, 1, 1, 1],
  [0, 0, 0, 0, 0, 0]
] → 6
*/
package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0},
	}
	fmt.Println(Calc(matrix))
}

func Calc(listShip [][]int) int {
	var res int
	for i := range listShip {
		for j := range listShip[i] {
			if listShip[i][j] == 1 {
				if j == 0 || listShip[i][j-1] == 0 {
					if i == 0 || listShip[i-1][j] == 0 {
						res++
					}
				}
			}
		}
	}
	return res
}
