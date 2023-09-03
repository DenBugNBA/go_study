package main

import "fmt"

func addToChan(i int, c chan int) {
	c <- i
}

func main() {
	c := make(chan int)

	go addToChan(1, c)
	go addToChan(2, c)
	go addToChan(3, c)
	go addToChan(4, c)
	go addToChan(5, c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
