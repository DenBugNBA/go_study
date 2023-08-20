package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// для простоты примера опускаем аргументы запроса и ответа
func DoHTTPCall() error {
	err := SendTCP()
	if err != nil {
		// оборачивается в виде "[название метода]: %w". %w — это плейсхолдер для ошибки
		return fmt.Errorf("send tcp: %w", err)
	}

	return nil
}

var errTCPConnectionIssue = errors.New("TCP connect issue")

func SendTCP() error {
	return errTCPConnectionIssue
}

func handleHTTPCall() error {
	err := DoHTTPCall()
	if err != nil {
		if errors.Is(err, errTCPConnectionIssue) {
			// в случае ошибки соединения ждем 1 секунду и пытаемся сделать запрос снова
			time.Sleep(1 * time.Second)
			return DoHTTPCall()
		}

		// обработка неизвестной ошибки
		log.Println("unknown error on HTTP call", err)
	}

	return nil
}

func main() {
	fmt.Println(handleHTTPCall()) // send tcp: TCP connect issue
}
