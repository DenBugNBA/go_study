package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	base, err := url.Parse("https://www.example.com")
	if err != nil {
		log.Println(err)
		return
	}

	// Добавляем путь к базовому URL, создавая "https://www.example.com/path"
	base.Path += "path"

	// Создаем query параметры запроса
	params := url.Values{}
	params.Add("id", "15")
	params.Add("name", "Dima")

	// Кодируем параметры запроса в строку и устанавливаем как часть запроса в URL
	base.RawQuery = params.Encode()

	// Выводим итоговый URL
	fmt.Printf("Encoded URL is %q\n", base.String())
}
