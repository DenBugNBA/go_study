package main

import (
	"fmt"
	"time"
)

type MyInt int

func (mi *MyInt) ShowByPointer() {
	fmt.Println(mi)
	fmt.Println(*mi)
}

func main() {
	ms := []MyInt{50, 60, 70, 80, 90}

	for _, m := range ms {
		go m.ShowByPointer()
	}

	/*
		only 90s:
		90
		90
		90
		90
		90
	*/

	time.Sleep(100 * time.Millisecond)
}
