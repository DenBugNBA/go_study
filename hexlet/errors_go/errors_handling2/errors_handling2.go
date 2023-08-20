package main

import (
	"errors"
	"log"
	"time"
)

// ошибка подключения к базе данных
type ConnectionErr struct{}

func (e ConnectionErr) Error() string {
	return "connection err"
}

// для простоты функция всегда возвращает ошибку подключения
func connectDB() error {
	return ConnectionErr{}
}

func main() {
	// цикл подключения к БД. Пытаемся 3 раза, если не удалось подсоединиться с первого раза.
	tries := 0
	for {
		if tries > 2 {
			log.Println("Can't connect to DB")
			break
		}

		err := connectDB()
		if err != nil {
			// если ошибка подключения, то ждем 1 секунду и пытаемся снова
			if errors.As(err, &ConnectionErr{}) {
				log.Println("Connection error. Trying to reconnect...")
				time.Sleep(1 * time.Second)
				tries++
				continue
			}

			// в противном случае ошибка критичная, логируем и выходим из цикла
			log.Println("connect DB critical error", err)
		}

		break
	}
}
