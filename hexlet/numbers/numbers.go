package main

import "fmt"

func main() {
	// cannot convert 5.05 (untyped float constant) to type int64
	//num := int64(5.05)

	num1 := 5.05
	fmt.Println(int(num1)) // 5

	num2 := int64(5.00)
	fmt.Println(num2) // 5

	numF := float64(5)
	fmt.Println(numF + 1.5) // 6.5
}
