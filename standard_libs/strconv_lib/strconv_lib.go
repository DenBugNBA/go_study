package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Itoa(-3)) // "-3"

	a, err := strconv.Atoi("-42")
	fmt.Println(a, err) // -42, <nil

	// ошибки компиляции нет, но значение всегда будет равно 0
	intNumber, err := strconv.Atoi("5.00")
	fmt.Println(intNumber, err) // 0, err

	intNumber, err = strconv.Atoi("Hello, World!")
	fmt.Println(intNumber, err) // 0, err

	stringWithF := "5.00"
	intNumber, err = strconv.Atoi(stringWithF)
	fmt.Println(intNumber, err) // 0, err

	s := "23.23456"
	// 2 параметр - bitSize - 32 или 64 (32 для float32, 64 для float64)
	// но нужно понимать, что метод все равно вернет float64
	floatNumber, err := strconv.ParseFloat(s, 64)
	fmt.Println(floatNumber, err) // 23.23456 <nil>

	// Конкретный пример для разных bitSize:
	s = "1.0000000012345678"
	result32, _ := strconv.ParseFloat(s, 32)
	result64, _ := strconv.ParseFloat(s, 64)
	fmt.Println("bitSize32:", result32) // вывод 1 (не уместились)
	fmt.Println("bitSize64:", result64) // вывод  1.0000000012345678
}
