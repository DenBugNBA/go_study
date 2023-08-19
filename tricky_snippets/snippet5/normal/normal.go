package main

import (
	"fmt"
)

type MyInt int

func (mi MyInt) Show() {
	fmt.Println(mi)
}

func (mi *MyInt) ShowByPointer() {
	fmt.Println(*mi)
}

func main() {
	ms := []MyInt{50, 60, 70, 80, 90}

	for _, m := range ms {
		m.ShowByPointer()
		m.Show()

		fmt.Println()
	}

	/*
		As expected:
		50
		50

		60
		60

		70
		70

		80
		80

		90
		90

	*/
}
