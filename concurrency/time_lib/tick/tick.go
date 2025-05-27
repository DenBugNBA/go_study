package main

import (
	"fmt"
	"time"
)

func main() {
	// не может быть остановлен и собран сборщиком мусора (до 1.23)
	// используйте если должен работать вечно
	c := time.Tick(time.Second)
	i := 0
	for tickTime := range c {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			break
		}
	}
}
