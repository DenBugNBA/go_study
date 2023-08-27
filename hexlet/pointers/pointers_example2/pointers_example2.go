package main

import "fmt"

type Creature struct {
	Species string
}

func main() {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println(*&creature) // shark

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)
	fmt.Println("*pointer =", *pointer)

	*pointer = "dog"
	fmt.Println(creature)
	fmt.Println(*pointer)

	var creature1 Creature = Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature1)
	changeCreature(creature1)
	fmt.Printf("3) %+v\n", creature1)

	fmt.Println()
	changeCreatureRef(&creature1)
	fmt.Printf("3) %+v\n", creature1)

	var creatureEmptyPointer *Creature
	fmt.Println(creatureEmptyPointer) // <nil>

	fmt.Println()

	// runtime error: invalid memory address or nil pointer dereference
	//changeCreatureRef(creatureEmptyPointer)

	fmt.Printf("1) %+v\n", creatureEmptyPointer)
	changeCreatureRefSmart(creatureEmptyPointer)
	fmt.Printf("3) %+v\n", creatureEmptyPointer)

	var creature2 *Creature
	creature2 = &Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature2)
	changeCreatureRef(creature2)
	fmt.Printf("3) %+v\n", creature2)

	creature3 := Creature{Species: "cat"}
	fmt.Println(creature3) // new
	//fmt.Println(creature3.String()) // new
	fmt.Println(creature3) // new

	creature3.Reset()
	fmt.Println(creature3) // new

	creature3.ResetRef()
	fmt.Println(creature3.Species)
}

func changeCreature(creature Creature) {
	creature.Species = "dog"
	fmt.Printf("2) %+v\n", creature)
}

func changeCreatureRef(creature *Creature) {
	creature.Species = "dog"
	fmt.Printf("2) %+v\n", creature)
}

func changeCreatureRefSmart(creature *Creature) {
	if creature == nil {
		fmt.Println("Nil pointer")
		return
	}
	creature.Species = "dog"
	fmt.Printf("2) %+v\n", creature)
}

// вызывается при fmt.Println(creature)
//func (c Creature) String() string {
//	c.Species = "new"
//	return c.Species
//}

func (c Creature) Reset() {
	c.Species = ""
}

func (c *Creature) ResetRef() {
	c.Species = "default"
}
