package main

import "fmt"

// gets a channel and prints the greeting by reading from channel
func greet(c <-chan string) {
	fmt.Println("Hello " + <-c + "!")
}

// gets a channels and writes a channel to it
func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func main() {
	fmt.Println("main() started")

	cc := make(chan chan string)

	go greeter(cc)

	c := <-cc

	go greet(c) // start `greet` goroutine using `c` channel

	// send data to `c` channel
	c <- "John"

	fmt.Println("main() stopped")
}
