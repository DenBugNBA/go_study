package main

import "fmt"

func main() {
	c := make(chan int, 1)      // здесь 1 - размер буфера
	fmt.Println(len(c), cap(c)) // 0 1
	c <- 5
	fmt.Println(len(c), cap(c)) // 1 1
	fmt.Println(<-c)            // 5
}
