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
}
