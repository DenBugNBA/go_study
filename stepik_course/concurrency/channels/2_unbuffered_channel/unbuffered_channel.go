package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Println(len(c), cap(c)) // 0 1

	// выполнение программы приведет к панике, так как
	// записанное значение одновременно с записью в канал не вычитывается
	// fatal error: all goroutines are asleep - deadlock!
	c <- 5
}
