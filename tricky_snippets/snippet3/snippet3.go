package main

import "fmt"

const (
	zero = iota
	one
	two
	three
)

const (
	a = iota + 1
	_
	b
	c
)

const (
	e = iota
	f = 42
	g = iota
	h
)

func main() {
	fmt.Println(zero, one, two, three) // 0 1 2 3

	fmt.Println(a, b, c) // 1 3 4

	fmt.Println(e, f, g, h) // 0 42 2 3
}
