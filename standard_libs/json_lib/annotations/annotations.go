package main

import (
	"encoding/json"
	"fmt"
)

// в общем виде аннотация выглядит так: `json:"используемое_имя,опция,опция"`
type myStruct struct {
	// при кодировании / декодировании будет использовано имя name, а не Name
	Name string `json:"name"`

	// при кодировании / декодировании будет использовано то же имя (Age),
	// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
	// то при кодировании оно будет опущено
	Age int `json:",omitempty"`

	// при кодировании / декодировании поле всегда игнорируется
	Status bool `json:"-"`
}

func main() {
	m := myStruct{Name: "John Connor", Age: 0, Status: true}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"name":"John Connor"}
}
