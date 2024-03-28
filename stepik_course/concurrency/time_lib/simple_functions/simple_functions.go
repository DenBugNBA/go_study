package main

import (
	"fmt"
	"time"
)

func main() {
	// программа засыпает на заданное время
	time.Sleep(time.Second * 2) // спим ровно 2 секунды

	// создает канал, который через заданное время вернет значение
	timer := time.After(time.Second)
	<-timer // значение будет получено из канала ровно через 1 секунду

	// создает канал, который будет посылать сигналы постоянно через заданный промежуток времени
	ticker := time.Tick(time.Second)
	count := 0

	for {
		<-ticker
		fmt.Println("очередной тик")
		count++
		if count == 3 {
			break
		}
	}
}
