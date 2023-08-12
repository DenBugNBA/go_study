package main

import (
	"fmt"
)

type User struct {
	email    string
	password string
}

// Функция получает структуру User по ссылке
func fillUserDataByValue(u User, email string, pass string) {
	u.email = email
	u.password = pass
}

// при объявлении указываем,
// что переменная должна быть указателем.
// Для этого ставим звездочку * перед типом данных
func fillUserData(u *User, email string, pass string) {
	u.email = email
	u.password = pass
}

func main() {
	u := User{}

	fillUserDataByValue(u, "test@test.com", "qwerty")
	fmt.Printf("no pointer:  %+v\n", u) // {email: password:}

	// передаем указатель с помощью амперсанда
	// & перед переменной
	fillUserData(&u, "test@test.com", "qwerty")
	fmt.Printf("points on func call %+v\n", u)
	// points on func call {email:test@test.com password:qwerty}

	// сразу инициализируем переменную с указателем
	userWithPointer := &User{}

	fillUserData(userWithPointer, "test@test.com", "qwerty")

	fmt.Printf("points on init %+v\n", userWithPointer)
	// points on init {email:test@test.com password:qwerty}

	fmt.Println()

	m := map[string]int{}

	// Мапы по умолчанию передаются с указателем:
	fillMap(m)

	fmt.Println(m) // map[random:1]
}

func fillMap(m map[string]int) {
	m["random"] = 1
}
