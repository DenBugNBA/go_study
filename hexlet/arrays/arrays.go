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

	fmt.Println()

	d := [3]int{1: 12}
	fmt.Println(d) // [0 12 0]

	fmt.Println()

	// сравнение массивов
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{3, 2, 1}

	fmt.Println(a == b) // true
	fmt.Println(a == c) // false

	fmt.Println()

	// пройтись по длине массива или среза, без элементов массива и индексов:
	for range c {
		fmt.Println("work")
	}

	fmt.Println()

	workArray := [3]int{1, 2, 3}
	workArray[0], workArray[1] = workArray[1], workArray[0]
	fmt.Println(workArray)
}
