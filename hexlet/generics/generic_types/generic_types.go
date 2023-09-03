package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	el1 := List[string]{val: "hello"}
	el2 := List[string]{next: &el1, val: "world"}

	fmt.Println(el2.next.val)
}
