package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK!"))
}

func main() {
	// создаем свой serverMux
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", handler)

	// Запускаем веб-сервер на порту 8080 с нашим serverMux (в прошлых примерах был nil)
	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
