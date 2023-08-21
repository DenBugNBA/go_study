package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	/*
		5
		5
		5
		5
		5
	*/

	time.Sleep(100 * time.Millisecond)
}
