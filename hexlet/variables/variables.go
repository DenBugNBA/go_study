package main

import "fmt"

func main() {
	var num1 int = 11
	fmt.Println(num1)

	// также можно объявить таким способом
	var num2 = 11
	fmt.Println(num2)

	// короткая запись
	num := 22
	fmt.Println(num)

	// двоеточие используется только при инициализации
	num = 33

	// переменные принято называть в camelCase:
	longTrickyName := "Josefina"
	fmt.Println(longTrickyName)

	var (
		a string // ""
		b bool   // false
		c int    // 0
	)

	fmt.Println(a, b, c)

	var i int
	j := i // j is an int

	i = 1
	fmt.Println(i, j) // 1, 0
}
