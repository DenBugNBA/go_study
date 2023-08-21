package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	/*
		4
		0
		1
		2
		3
	*/

	time.Sleep(100 * time.Millisecond)
}
