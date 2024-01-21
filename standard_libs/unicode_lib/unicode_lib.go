package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.Is(unicode.Cyrillic, 'ы')) // true
	fmt.Println(unicode.Is(unicode.Latin, 'ы'))    // false

	fmt.Println(unicode.IsDigit('1')) // true

	fmt.Println(unicode.IsLetter('a')) // true

	fmt.Println(unicode.IsLower('A')) // false

	fmt.Println(unicode.IsSpace('\t')) // true

	fmt.Println(unicode.IsUpper('A')) // true

	fmt.Println(string(unicode.ToLower('F'))) // f

	fmt.Println(string(unicode.ToUpper('f'))) // F
}
