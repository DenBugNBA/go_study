package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	for elem := range c {
		fmt.Println(elem)
	}
	fmt.Println("done")
}
