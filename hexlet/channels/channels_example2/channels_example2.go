package main

import (
	"fmt"
	"time"
)

func main() {
	// создаем канал, в который будем отправлять сообщения
	msgCh := make(chan string)

	// вызываем функцию асинхронно в горутине
	go printer(msgCh)

	msgCh <- "hello"
	msgCh <- "concurrent"
	msgCh <- "world"

	// закрываем канал
	close(msgCh)

	// и ждем, пока printer закончит работу
	time.Sleep(100 * time.Millisecond)
}

func printer(msgCh chan string) {
	// читаем из канала, пока он открыт
	for msg := range msgCh {
		fmt.Println(msg)
	}

	fmt.Println("printer has finished")
}
