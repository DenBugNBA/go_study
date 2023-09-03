package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			message <- strconv.Itoa(i)
			time.Sleep(time.Millisecond * 500)
		}

		close(message)
	}()

	for {
		msg, open := <-message
		if !open {
			break
		}

		fmt.Println(msg)
	}

	// та же запись, что выше
	for msg := range message {
		fmt.Println(msg)
	}
}
