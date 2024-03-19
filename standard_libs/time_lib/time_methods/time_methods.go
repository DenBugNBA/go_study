package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// (year int, month Month, day int)
	fmt.Println(current.Date()) // 2020 May 15

	fmt.Println(current.Year()) // 2020

	fmt.Println(current.Month()) // May

	fmt.Println(current.Day()) // 15

	// (hour int, min int, sec int)
	fmt.Println(current.Clock()) // 17 45 12

	fmt.Println(current.Hour()) // 17

	fmt.Println(current.Minute()) // 45

	fmt.Println(current.Second()) // 12

	fmt.Println(current.Unix()) // 1589546712

	fmt.Println(current.Weekday()) // Friday

	fmt.Println(current.YearDay()) // 136
}
