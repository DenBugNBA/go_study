package main

import "fmt"

func main() {
	x := 10

	switch x {
	// default всегда выполняется последним независимо от расположения в конструкции
	default:
		fmt.Println("default case")
	case 10:
		fmt.Println("case 10")
	} // case 10

	fmt.Println()

	// выражение отсутствует. Для компилятора выглядит как: switch true
	switch {
	default:
		fmt.Println("default case")
	case x == 10:
		fmt.Println("equal 10 case")
		fallthrough
	case x <= 10:
		fmt.Println("less or equal 10 case")
	}
	// equal 10 case
	// less or equal 10 case

	fmt.Println()

	x = 5

	switch {
	case x == 5:
		fmt.Println("Bob")
		fallthrough
	case x <= 5:
		fmt.Println("Alice")
		fallthrough
	case x > 55:
		fmt.Println("Joe")
		fallthrough
	default:
		fmt.Println("Default")
	}
}
