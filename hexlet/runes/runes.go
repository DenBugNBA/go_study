package main

import (
	"fmt"
	"unicode"
)

func IsASCIITeacher(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}

	return true
}

func IsASCII(s string) bool {
	return len([]rune(s)) == len([]byte(s))
}

func main() {
	emoji := []rune("привет😀")

	fmt.Println(len(emoji)) // 7

	for i := 0; i < len(emoji); i++ {
		fmt.Println(emoji[i], string(emoji[i])) // выводим код символа и его строковое представление
	}

	fmt.Println()

	s := "hey😉"

	// cannot convert ([]byte)(s) (type []byte) to type []rune
	// rs := []rune([]byte(s))

	// cannot convert ([]rune)(s) (type []rune) to type []byte
	// bs := []byte([]rune(s))

	rs := []rune(s)
	fmt.Println(rs)

	bs := []byte(s)
	fmt.Println(bs)

	fmt.Println()

	emoji2 := []rune("cool😀")

	for _, ch := range emoji2 {
		fmt.Println(ch, string(ch)) // выводим код символа и его строковое представление
	}

	fmt.Println()

	/*
		В Go присутствует синтаксический сахар при обходе строки. Если использовать конструкцию
		for range, строка автоматически будет преобразована в []rune, то есть обход будет по Юникод символам:
	*/
	unicodeString := "привет"

	for _, ch := range unicodeString {
		fmt.Println(ch, string(ch)) // выводим код символа и его строковое представление
	}
}
