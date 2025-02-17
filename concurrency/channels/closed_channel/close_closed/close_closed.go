package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	fmt.Println("closed")
	// panic: close of closed channel
	close(ch)
	fmt.Println("closed twice")
}
