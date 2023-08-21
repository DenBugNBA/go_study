package main

import (
	"fmt"
	"time"
)

func main() {
	// выведет сообщение в горутине
	go fmt.Println("Hello concurrent world")

	// если не подождать, то программа закончится, не успев вывести сообщение
	time.Sleep(100 * time.Millisecond)
}
