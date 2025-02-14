package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	var c chan int
	fmt.Println(c) // <nil>

	fmt.Println()

	simpleChan := make(chan int)

	fmt.Printf("type of `c` is %T\n", simpleChan)
	fmt.Printf("value of `c` is %v\n", simpleChan)

	fmt.Println()

	chanForGreet := make(chan string)
	go greet(chanForGreet)
	chanForGreet <- "John"

	fmt.Println("main() stopped")
}
