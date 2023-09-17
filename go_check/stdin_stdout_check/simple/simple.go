package main

import "fmt"

func main() {
	var num int

	fmt.Println(num)  // 0
	fmt.Println(&num) // 0xc0000120d8

	fmt.Scan(&num)

	fmt.Println(num)
}
