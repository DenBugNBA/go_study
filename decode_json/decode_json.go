package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Message struct {
	Sender string `json:"sender"` // ставим тег с описанием JSON поля
	Text   string `json:"text"`
}

// инициализация ошибки через конструктор стандартного пакета errors
var errEmptyMessage = errors.New("empty message")

// возвращаем ошибку в случае неожиданного поведения
func DecodeJSON(rawMsg string) (Message, error) {
	// если нам передали пустую строку, возвращаем ошибку об этом
	if len(rawMsg) == 0 {
		return Message{}, errEmptyMessage
	}

	msg := Message{}

	// декодируем строку в структуру
	err := json.Unmarshal([]byte(rawMsg), &msg)
	if err != nil {
		return Message{}, fmt.Errorf("unmarshal: %w", err)
	}

	return msg, nil
}

func main() {
	msg, err := DecodeJSON("")
	if errors.Is(err, errEmptyMessage) {
		// { } empty message
		fmt.Println(msg, err)
	}

	msg, err = DecodeJSON("hello")
	if err != nil {
		// { } unmarshal: invalid character 'h' looking for beginning of value
		fmt.Println(msg, err)
	}

	msg, err = DecodeJSON(`{"sender":"hexlet","text":"Go,Go,Go"}`)
	// {hexlet Go,Go,Go} <nil>
	fmt.Println(msg, err)
}
