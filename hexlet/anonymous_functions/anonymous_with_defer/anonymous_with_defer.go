package main

import "fmt"

func ExampleDefer1() {
	defer func() { fmt.Println(1) }()

	defer func() { fmt.Println(2) }()

	defer func() { fmt.Println(3) }()
}

func main() {
	ExampleDefer1()
	// 3
	// 2
	// 1
}
