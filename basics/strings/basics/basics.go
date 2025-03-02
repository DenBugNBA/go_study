package main

import "fmt"

func main() {
	// Строки в Go — это иммутабельные массивы байт
	s := "hey"

	// Байты представляют ASCII символы, а в кодовой таблице ASCII символов 256 кодов:
	fmt.Println(s[0], s[1], s[2])                         // 104 101 121
	fmt.Println(string(s[0]), string(s[1]), string(s[2])) // h e y

	bs := []byte(s)
	fmt.Println(bs)         // [104 101 121]
	fmt.Println(string(bs)) // hey

	fmt.Println()

	wordRussian := "привет"

	// len считает кол-во байт, а не кол-во символов
	fmt.Println(len(wordRussian), "- len(wordRussian)")

	q := `
		SELECT *
		FROM person
		WHERE age > 18
	`
	fmt.Println(q)
}
