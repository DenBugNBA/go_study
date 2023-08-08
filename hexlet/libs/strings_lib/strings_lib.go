package main

import (
	"fmt"
	"strings"
	"unicode"
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

func LatinLetters(s string) string {
	sb := &strings.Builder{}

	for _, r := range s {
		// которая проверяет, что руна является латинским символом или нет
		if unicode.Is(unicode.Latin, r) {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func main() {
	fmt.Println(strings.Join([]string{"Hello", "World"}, " "))                                         // Hello World
	fmt.Println(strings.Trim(strings.Join([]string{"Hello", "World", ""}, " "), " ") == "Hello World") // true

	fmt.Println(strings.Split("Hello World", " ")) // [Hello World]

	fmt.Println(strings.Fields("  one    two\tthree four ")) // [one two three four]

	fmt.Println(strings.Trim(" hello ", " "))            // "hello"
	fmt.Println(strings.TrimSpace(" hello ") == "hello") // true

	fmt.Println(strings.ToLower("HELLO")) // "hello"
	fmt.Println(strings.ToUpper("hello")) // "HELLO"

	fmt.Println(strings.HasPrefix("Mr. Smith", "Mr.")) // true

	fmt.Println(strings.Contains("hello", "ll")) // true

	fmt.Println(strings.ReplaceAll("one two three", " ", "_"))

	fmt.Println(Greetings("  ИВАН  "))
	fmt.Println(ModifySpaces("Hello World!", "underscore"))

	sb := strings.Builder{}

	sb.WriteString("hello")
	sb.WriteString(" ")
	sb.WriteString("world")

	fmt.Println(sb.String()) // "hello world"
}
