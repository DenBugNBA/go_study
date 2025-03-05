package main

import "fmt"

func change(a *int) {
	t := 2

	a = &t

	fmt.Println(a == &t)

	*a = 10

	fmt.Println("t =", t)
}

func main() {
	a := 3
	change(&a)

	fmt.Println(a) // 3
}
