package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Привет, я анонимная горутина")
	}()
	time.Sleep(1 * time.Second)
}
