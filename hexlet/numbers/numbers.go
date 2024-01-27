package main

import "fmt"

func main() {
	// cannot convert 5.05 (untyped float constant) to type int
	// num := int(5.05)

	num1 := 5.05
	fmt.Println(int(num1)) // 5

	num2 := int(5.00)
	fmt.Println(num2) // 5

	// num := 2
	// 4.5 (untyped float constant) truncated to int
	// fmt.Println(num + 4.5)

	// превышено максимальное значение для данного типа данных
	var big int64 = 129
	var little = int8(big)
	fmt.Println(little) // -127

	a := 5.0 / 2
	fmt.Println(a) // 2.5

	negativeNum := -5
	// ошибки компиляции нет, но число было преобразовано в положительное путем прибавления MAX_UINT+1.
	// MAX_UINT = 18446744073709551615 (+ 1) - 5 =
	x := uint(negativeNum) // 18446744073709551611
	fmt.Println(x)

	// cannot convert -42 (untyped int constant) to type uint
	// x = uint(-42)
}
