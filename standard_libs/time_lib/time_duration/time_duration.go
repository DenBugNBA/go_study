package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	past := now.AddDate(0, 0, -1)
	future := now.AddDate(0, 0, 1)

	// вычисляет период между текущим моментом и заданным временем в прошлом
	fmt.Println(time.Since(past).Round(time.Second)) // 24h0m0s

	// вычисляет период между текущим моментом и заданным временем в будущем
	fmt.Println(time.Until(future).Round(time.Second)) // 24h0m0s

	// преобразует строку в Duration с использованием аннотаций:
	// "ns" - наносекунды,
	// "us" - микросекунды,
	// "ms" - миллисекунды,
	// "s" - секунды,
	// "m" - минуты,
	// "h" - часы.
	dur, err := time.ParseDuration("1h12m3s")
	if err != nil {
		panic(err)
	}
	fmt.Println(dur.Round(time.Hour).Hours()) // 1

}
