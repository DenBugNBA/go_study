package main

import (
	"fmt"
	"strconv"
)

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func main() {
	fmt.Println(IntToString(-3))

	a, err := strconv.Atoi("-42")
	fmt.Println(a, err) // -42, <nil>

	// ошибки компиляции нет, но число было преобразовано в положительное путем прибавления MAX_UINT+1. MAX_UINT = 18446744073709551615
	x := uint(a) // 18446744073709551574
	fmt.Println(x)

	// error: cannot convert -42 (untyped int constant) to type uint
	//x = uint(-42)

	// ошибки компиляции нет, но значение всегда будет равно 0
	intNumber, err := strconv.Atoi("5.00")
	fmt.Println(intNumber, err) // 0, err

	intNumber, err = strconv.Atoi("Hello, World!")
	fmt.Println(intNumber, err) // 0, err

	stringWithF := "5.00"
	intNumber, err = strconv.Atoi(stringWithF)
	fmt.Println(intNumber, err) // 0, err
}
