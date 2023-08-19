package main

import (
	"fmt"
	"os"
)

type user struct {
	name string
	age  int
}

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()

	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false

	fmt.Println()

	var pointerInt *int = nil

	fmt.Println(pointerInt)        // <nil>
	fmt.Println(pointerInt == nil) // true

	fmt.Println()

	var pointerStruct *user = nil

	fmt.Println(pointerStruct)        // <nil>
	fmt.Println(pointerStruct == nil) // true
}
