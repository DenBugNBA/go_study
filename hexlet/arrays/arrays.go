package main

import "fmt"

func main() {
	// можно опустить размер массива и автоматически определить его
	var fruits = [...]string{"apple", "banana", "orange"}
	fmt.Println(fruits)      // [apple banana orange]
	fmt.Println(len(fruits)) // 3

	fmt.Println()

	// автоматически заполняется значениями по умолчанию для своего типа данных
	var nums [5]int
	fmt.Println(nums) // [0 0 0 0 0]
}
