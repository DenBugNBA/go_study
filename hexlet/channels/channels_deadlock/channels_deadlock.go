package main

import (
	"fmt"
)

func main() {
	numCh := make(chan int)

	<-numCh // программа зависнет здесь и будет ошибка: fatal error: all goroutines are asleep - deadlock!

	fmt.Println("program has ended") // эта строка никогда не выведется
}
