package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	// проверяем что метод POST
	if r.Method != "POST" {
		w.Write([]byte("Разрешен только метод POST!"))
		return
	}

	// читаем входящее тело запроса
	bytesBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Плохое тело запроса"))
		return
	}

	// печатаем тело запроса как строку
	fmt.Println(string(bytesBody))

	// отвечаем клиенту, что все хорошо
	w.Write([]byte("OK!"))

	return
}

func main() {
	// Регистрируем обработчик для пути "/body"
	http.HandleFunc("/body", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
