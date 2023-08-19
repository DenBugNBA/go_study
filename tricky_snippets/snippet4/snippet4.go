package main

import "fmt"

func main() {
	a := 2

	func() {
		a := a
		fmt.Println(a) // 2

		a++
		fmt.Println(a) // 3
	}()

	fmt.Println(a) // 2
}
