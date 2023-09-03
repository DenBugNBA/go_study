package main

import (
	"fmt"
)

func main() {
	numCh := make(chan int)

	// fatal error: all goroutines are asleep - deadlock!
	<-numCh // программа зависнет здесь и будет ошибка

	fmt.Println("program has ended") // эта строка никогда не выведется
}
