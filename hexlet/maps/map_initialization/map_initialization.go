package main

import "fmt"

func main() {
	// map не инициализирована
	var users map[string]int
	fmt.Println(users)

	// panic: assignment to entry in nil map
	// users["John"] = 10

	// инициализация с помощью встроенной функции make:
	m1 := make(map[int]int)

	// инициализация с помощью использования литерала отображения:
	m2 := map[int]int{
		// Пары ключ:значение указываются при необходимости
		12: 2,
		1:  5,
	}

	fmt.Println(m1) // map[]
	fmt.Println(m2) // map[1:5 12:2]
}
