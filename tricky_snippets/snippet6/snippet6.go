package main

import "fmt"

type I interface {
	A()
}

type S struct{}

func (s *S) A() {}

func main() {
	var s *S = nil
	fmt.Println(s == nil) // true
	fmt.Println(s)        // <nil>

	fmt.Println()

	//структура указателя на nil
	var i I = s
	fmt.Println(i == nil) // false
	fmt.Println(i)        // <nil>
}
