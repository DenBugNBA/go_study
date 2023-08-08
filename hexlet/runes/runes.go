package main

import "fmt"

func main() {
	emoji := []rune("привет😀")

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
}
