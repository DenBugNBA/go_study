package main

import "fmt"

var counter = 0 //  общий ресурс

func main() {
	ch := make(chan bool) // канал

	for i := 1; i < 5; i++ {
		go work(i, ch)
	}

	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The End")
}

func work(number int, ch chan bool) {
	counter = 0

	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}

	ch <- true
}
