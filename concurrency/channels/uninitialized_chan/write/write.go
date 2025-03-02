package main

import "fmt"

func main() {
	var ch chan int
	ch <- 1 // deadlock - chan send (nil chan)
	fmt.Println(<-ch)
}
