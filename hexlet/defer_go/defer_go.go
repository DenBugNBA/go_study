package main

import (
	"fmt"
)

func main() {
	// stack
	defer fmt.Println("3rd")
	defer fmt.Println("2nd")

	fmt.Println("1st")
}
