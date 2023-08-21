package main

import (
	"encoding/json"
	"fmt"
)

// структура публична
type Person struct {
	// поле публично
	Name string

	// поле приватно: можно обращаться только внутри текущего пакета
	wallet string
}

type User struct {
	ID        int64  `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
}

type UserJSON struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}

type UserValidate struct {
	ID        int64  `validate:"required"`
	Email     string `validate:"required,email"`
	FirstName string `validate:"required"`
}

func main() {
	p := Person{wallet: "rub"}
	fmt.Println(p.wallet) // rub

	u := UserJSON{}

	fmt.Println(u.ID) // 0

	emptyUser, _ := json.Marshal(u)
	fmt.Println(emptyUser)
	fmt.Println(string(emptyUser))

	u.ID = 22
	u.Email = "test@test.com"
	u.FirstName = "John"

	bs, _ := json.Marshal(u)

	fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}
}
