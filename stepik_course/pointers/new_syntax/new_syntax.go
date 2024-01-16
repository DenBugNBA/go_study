package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	fmt.Println(xPtr)  // 0xc0000120d8
	fmt.Println(*xPtr) // 0
	one(xPtr)
	fmt.Println(*xPtr) // 1
}
