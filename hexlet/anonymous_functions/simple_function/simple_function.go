package main

import "fmt"

// Когда анонимная функция использует переменные, объявленные за ее рамками,
// ее называют замыканием.
func externalFunction() func() {
	text := "TEXT"

	return func() {
		fmt.Println(text)
	}
}

func ExampleEnvironment() {
	fn := externalFunction()
	fn()
}

func main() {
	ExampleEnvironment() // TEXT
}
