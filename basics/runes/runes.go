package main

import "fmt"

func IsASCII(s string) bool {
	return len([]rune(s)) == len([]byte(s))
}

func main() {
	simpleRune := 'r'                           // int32
	fmt.Println(simpleRune, string(simpleRune)) // 114 r

	fmt.Println()

	emoji := []rune("привет😀")
	fmt.Println("len:", len(emoji)) // 7

	for i := 0; i < len(emoji); i++ {
		fmt.Println(emoji[i], string(emoji[i])) // выводим код символа и его строковое представление
	}

	fmt.Println()

	s := "hey😉"

	rs := []rune(s)
	fmt.Println(rs) // [104 101 121 128521]

	bs := []byte(s)
	fmt.Println(bs) // [104 101 121 240 159 152 137]
	fmt.Println(string(bs)) // hey😉

	fmt.Println()

	unicodeString := "привет"
	// если использовать конструкцию for range, строка автоматически будет преобразована в []rune,
	// то есть обход будет по Юникод символам:
	for i, ch := range unicodeString {
		fmt.Println(i, ch, string(ch))
	}

	fmt.Println()

	unicodeString = "hello"
	for i, ch := range unicodeString {
		fmt.Println(i, ch, string(ch)) // i - индекс стартового байта символа
	}
}
