package main

import "fmt"

func main() {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println(*&creature) // shark

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	*pointer = "dog"
	fmt.Println(creature)
}
