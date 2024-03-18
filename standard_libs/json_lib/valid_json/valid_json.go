package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func main() {
	m := user{Name: "John Connor", Email: "test email"}

	data, _ := json.Marshal(m)

	data = bytes.Trim(data, "{") // испортим json удалив '{'

	// функция json.Valid возвращает bool, true - если json правильный
	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}

	fmt.Printf("%s", data) // вывод: "Name":"John Connor","Email":"test email","Status":false,"Language":null}
}
