package main

import "fmt"

func zeroval(iVal int) {
	iVal = 0
}

func zeroptr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // 1

	zeroval(i)
	fmt.Println("zeroval:", i) // 1

	zeroptr(&i)
	fmt.Println("zeroptr:", i) // 0

	fmt.Println("pointer:", &i) // 0xc0000120d8

	a := 1
	pointer := &a
	*pointer = 2
	fmt.Println(a) // 2

	nums := []int{1, 2, 3}
	numsPointer := &nums

	fmt.Println(*numsPointer)
}
