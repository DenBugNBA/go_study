package main

import "fmt"

func main() {
	// дробная часть отбрасывается
	var a int = 11 / 6
	fmt.Println(a) // 1

	// Чтобы получить в результате деления вещественное число,
	// как минимум один из операндов также должен представлять собой вещественное число
	// и результат мы должны при этом тоже сохранять в переменную вещественного типа:
	var m float32 = 10.0 / 6
	fmt.Println(m) // 1.6666666
}
