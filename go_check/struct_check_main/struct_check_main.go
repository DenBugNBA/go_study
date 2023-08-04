package main

import (
	"fmt"
	"go_study/go_check/struct_check"
)

func main() {
	// Cannot assign a value to the unexported field '_Age'
	p1 := struct_check.Person{Name: "Jack"}
	fmt.Println(p1)
}
