package main

import "fmt"

func main() {
	// при выводе нескольких объектов вставляет между ними пробелы, если среди них нет строк.
	fmt.Print("ef", "ef") // efef
	fmt.Println()

	fmt.Print("Ivan", 27) // Ivan27
	fmt.Println()

	fmt.Print(33, 27) // 33 27
	fmt.Println()

	// всегда ставит пробелы между выводимыми объектами, плюс добавляет новую строку.
	fmt.Println("ef", "ef") // ef ef
}
