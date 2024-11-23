package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Task - структура для представления объекта Task
type Task struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Создаем экземпляр структуры Task
	task := Task{
		UserID:    1,
		ID:        2,
		Title:     "наш title",
		Completed: true,
	}

	// Кодируем структуру Task в формат JSON
	jsonReq, err := json.Marshal(task)
	if err != nil {
		log.Println(err)
		return
	}

	// URL сервера
	baseURL := "https://jsonplaceholder.typicode.com/posts"

	// Создаем новый HTTP-запрос с методом POST
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Println("Ошибка при создании запроса:", err)
		return
	}

	// Устанавливаем заголовки запроса
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Отправляем запрос
	client := &http.Client{}        // создаем http клиент
	response, err := client.Do(req) // передаем выше подготовленный запрос на отправку
	if err != nil {
		log.Println("Ошибка при выполнении запроса: ", err)
		return
	}

	defer response.Body.Close() // не забываем закрыть тело

	// Вывод статуса ответа
	fmt.Println("Статус ответа:", response.Status)

	// Читаем и конвертируем тело ответа в байты
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	// Конвертируем тело ответа в строку и выводим
	bodyString := string(bodyBytes)
	fmt.Printf("API ответ в форме строки: \n%s\n", bodyString)

	// Конвертируем тело ответа в Task struct
	var taskStruct Task
	err = json.Unmarshal(bodyBytes, &taskStruct)
	if err != nil {
		log.Println(err)
	}

	// Выводим структуру Task
	fmt.Printf("API ответ в форме struct:\n%+v\n", taskStruct)
}
