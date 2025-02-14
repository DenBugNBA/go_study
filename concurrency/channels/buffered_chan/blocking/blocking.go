package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	c <- 3
	c <- 4 // blocks here

	fmt.Println("main() stopped")
}
