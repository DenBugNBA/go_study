package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step:", i, "time:", tickTime)
		if i >= 5 {
			// надо останавливать, иначе потечет (до 1.23)
			ticker.Stop()
			break
		}
	}
	fmt.Println("total:", i)
}
