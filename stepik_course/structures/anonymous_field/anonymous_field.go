package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	// отношение андроид «является» личностью
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Name = "John"
	a.Talk()
}
