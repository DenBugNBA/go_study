package main

import (
	"fmt"
	"strings"
)

func Greetings(name string) string {
	return fmt.Sprintf("Привет, %s!", strings.Title(strings.Trim(strings.ToLower(name), " ")))
}

func ModifySpaces(s, mode string) string {
	var replaceChar string

	switch mode {
	case "dash":
		replaceChar = "-"
	case "underscore":
		replaceChar = "_"
	default:
		replaceChar = "*"
	}

	return strings.ReplaceAll(s, " ", replaceChar)
}

func main() {
	fmt.Println(strings.Join([]string{"Hello", "World"}, " "))                                         // Hello World
	fmt.Println(strings.Trim(strings.Join([]string{"Hello", "World", ""}, " "), " ") == "Hello World") // true

	fmt.Println(strings.Split("Hello World", " ")) // [Hello World]

	fmt.Println(strings.Trim(" hello ", " ")) // "hello"

	fmt.Println(strings.ToLower("HELLO")) // "hello"
	fmt.Println(strings.ToUpper("hello")) // "HELLO"

	fmt.Println(strings.HasPrefix("Mr. Smith", "Mr.")) // true

	fmt.Println(strings.Contains("hello", "ll")) // true

	fmt.Println(strings.ReplaceAll("one two three", " ", "_"))

	fmt.Println(Greetings("  ИВАН  "))
	fmt.Println(ModifySpaces("Hello World!", "underscore"))

}
