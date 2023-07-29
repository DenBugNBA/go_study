package main

import (
	"fmt"
	"strings"
)

func Greetings(name string) string {
	return fmt.Sprintf("Привет, %s!", strings.Title(strings.Trim(strings.ToLower(name), " ")))
}

func main() {
	fmt.Println(strings.Trim(" hello ", " ")) // "hello"
	fmt.Println(strings.ToLower("HELLO"))     // "hello"
	fmt.Println(strings.ToTitle("HELLO"))     // "hello"

	fmt.Println(Greetings("  ИВАН  "))
}
