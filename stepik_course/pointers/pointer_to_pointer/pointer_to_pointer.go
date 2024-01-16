package main

import "fmt"

func main() {
	a := 200
	b := &a
	*b++

	// c -указатель на указатель
	c := &b
	fmt.Println(c)  // 0xc00009a018
	fmt.Println(*c) // 0xc0000a2000 - адрес из b

	**c++
	fmt.Println(a)
}
