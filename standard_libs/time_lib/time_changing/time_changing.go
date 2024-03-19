package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// изменяет дату в соответствии с параметром - "продолжительностью"
	future := now.Add(time.Hour * 12) // перемещаемся на 12 часов вперед
	fmt.Println(future)

	// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
	past := now.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

	// вычисляет время, прошедшее между двумя датами
	fmt.Println(future.Sub(past)) // 10332h0m0s
}
