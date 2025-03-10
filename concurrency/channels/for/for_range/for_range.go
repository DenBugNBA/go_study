package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i
	}
	close(c) // close channel
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("main() stopped")
}
