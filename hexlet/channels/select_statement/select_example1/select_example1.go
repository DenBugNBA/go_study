package main

import (
	"fmt"
	"time"
)

func main() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			message1 <- "Прошло пол секунды"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			message2 <- "Прошло 2 секунды"
		}
	}()

	for {
		select {
		case msg1 := <-message1:
			fmt.Println(msg1)

		case msg2 := <-message2:
			fmt.Println(msg2)

			//fmt.Println(<-message1)
			//fmt.Println(<-message2)
		}
	}
}
