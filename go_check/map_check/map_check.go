package main

import "fmt"

func main() {
	nameToAge := make(map[string]int)

	nameToAge["Jack"] = 1

	fmt.Println(nameToAge["Jack"]) // 1
	fmt.Println(nameToAge["jack"]) // 0

	value, exists := nameToAge["Nick"]
	if exists {
		fmt.Println(value)
	}
}
