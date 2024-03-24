package main

import "fmt"

func main() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go myFunc(c)
	}
	fmt.Println(<-c) // 5
}

func myFunc(c chan int) {
	c <- cap(c)
}
