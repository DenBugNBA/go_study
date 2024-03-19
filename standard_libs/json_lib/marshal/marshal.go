package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name         string
	Age          int
	Status       bool
	Values       []int
	privateNotes string // неэкспортируемые поля игнорируются
}

func main() {
	s := myStruct{
		Name:         "John Connor",
		Age:          35,
		Status:       true,
		Values:       []int{15, 11, 37},
		privateNotes: "something",
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
}
