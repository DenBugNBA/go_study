package main

import (
	"errors"
	"fmt"
)

func multiplyAny(numbers ...int) int {
	result := 1

	for _, number := range numbers {
		result *= number
	}

	return result
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide on zero")
	}

	return x / y, nil
}

func main() {
	fmt.Println(multiply(2, 3)) // 6

	res, err := divide(5, 2)
	fmt.Println(res, err) // 2 <nil>

	res, err = divide(5, 0)
	fmt.Println(res, err) // 0 cannot divide on zero

	fmt.Println(multiplyAny(1, 2, 3))    // 6
	fmt.Println(multiplyAny(1, 2, 3, 4)) // 24

	// empty() (no value) used as value
	// emptyVal := empty()
}

func empty() {}
