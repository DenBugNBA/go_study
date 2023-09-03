package main

import "fmt"

func main() {
	strCh := make(chan string)

	// fatal error: all goroutines are asleep - deadlock!
	// strCh <- "hello"

	go readChannel(strCh)

	strCh <- "hello"

	// fatal error: all goroutines are asleep - deadlock!
	// strCh <- "hello"
}

func readChannel(strCh chan string) {
	fmt.Println(<-strCh)
}
