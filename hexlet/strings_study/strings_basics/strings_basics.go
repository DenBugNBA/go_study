package main

import (
	"fmt"
	"reflect"
)

func NextASCII(b byte) byte {
	return b + 1
}

func PrevASCII(b byte) byte {
	return b - 1
}

func ShiftASCII(s string, step int) string {
	resultBytes := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		resultBytes = append(resultBytes, s[i]+uint8(step))
	}

	return string(resultBytes)
}

func main() {
	// Строки в Go — это иммутабельные массивы байт
	s := "hey"

	q := `
		SELECT *
		FROM person
		WHERE age > 18
	`
	fmt.Println(q)

	// Байты представляют ASCII символы, а в кодовой таблице ASCII символов 256 кодов:
	fmt.Println(s[0], s[1], s[2]) // 104 101 121

	fmt.Println(string(s[0]), string(s[1]), string(s[2])) // h e y

	bs := []byte(s)
	fmt.Println(bs) // [104 101 121]

	fmt.Println(string(bs)) // hey

	asciiCh := byte('Z')
	asciiChStr := string(asciiCh)

	fmt.Println(reflect.TypeOf(asciiCh), asciiCh)       // uint8 90
	fmt.Println(reflect.TypeOf(asciiChStr), asciiChStr) // string Z

	fmt.Println(NextASCII(byte('a'))) // 98
	fmt.Println(PrevASCII(byte('b'))) // 97

	// Таким способом можно обойти только строки, состоящие из ASCII символов.
	word := "hello"
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}

	fmt.Println()

	// Если строка содержит мультибайтовые символы, вывод будет некорректен:
	wordRussian := "привет"
	fmt.Println(wordRussian, "- wordRussian")

	// Функция считает кол-во байт, а не кол-во символов
	fmt.Println(len(wordRussian), "- len(wordRussian)")

	for i := 0; i < len(wordRussian); i++ {
		fmt.Println(wordRussian[i])
		fmt.Println(string(wordRussian[i]))
	}
	// Go будет "делить" символы, кодируемые несколькими байтами на одиночные, и переводить по таблице ASCII

	fmt.Println(ShiftASCII("abc", 0))  // "abc"
	fmt.Println(ShiftASCII("abc1", 1)) // "bcd2"
}
