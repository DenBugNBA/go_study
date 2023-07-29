package main

import (
	"errors"
	"fmt"
)

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
	res, err := divide(5, 2)
	fmt.Println(res, err)

	res, err = divide(5, 0)
	fmt.Println(res, err)
}
