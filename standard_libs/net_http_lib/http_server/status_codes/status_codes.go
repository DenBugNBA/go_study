package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет!"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed) // вернем 405
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет 2!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/2", handler2)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
