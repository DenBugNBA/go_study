package main

import "fmt"

func main() {
	a := 1

	//	The deferred call's arguments are evaluated immediately,
	//	but the function call is not executed until the surrounding function returns.
	defer fmt.Println(a)

	a = 2
	fmt.Println(a)
	// 2
	// 1
}
