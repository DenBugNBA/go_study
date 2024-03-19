package main

import (
	"fmt"
	"time"
)

func main() {
	firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	// true если позже
	fmt.Println(firstTime.After(secondTime)) // true

	// true если раньше
	fmt.Println(firstTime.Before(secondTime)) // false

	// true если равны
	fmt.Println(firstTime.Equal(secondTime)) // false
}
