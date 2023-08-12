package main

import (
	"fmt"
	"go_study/go_check/struct_check"
)

func main() {
	// Cannot assign a value to the unexported field '_Age'
	// p1 := struct_check.Person{Name: "Jack", _Age: 10}

	p1 := struct_check.Person{Name: "Jack"}
	fmt.Println(p1)

	// Fields are assigned without explicit names
	// car := struct_check.Car{"Lada", 10}
	car := struct_check.Car{Name: "Lada", Speed: 10}
	fmt.Println(car)
}
