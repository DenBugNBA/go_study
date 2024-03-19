package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow, "- time now")

	fmt.Println(timeNow.Format("02-01-2006 15:04:05")) // 15-05-2020 09:58:16

	timeNowUnix := timeNow.Unix() // number of seconds elapsed since January 1, 1970, UTC.
	fmt.Println(timeNowUnix, "- seconds now")

	fmt.Println()

	const minSeconds = 94608000

	fmt.Println(time.Unix(minSeconds, 0), "- min seconds")

	fmt.Println()

	randomSeconds := rand.Int63n(timeNowUnix-minSeconds) + minSeconds
	fmt.Println(randomSeconds, "- random seconds")

	randomTime := time.Unix(randomSeconds, 0)
	fmt.Println(randomTime, "- random time")

	fmt.Println(randomTime.String())

	fmt.Println()

	// возвращает дату и время в соответствии с заданными параметрами
	someTime := time.Date(
		2020,
		time.May,
		15,
		10,
		13,
		12,
		45,
		time.UTC, // временная зона
	)
	fmt.Println(someTime, "- some time")

	// возвращает дату и время в соответствии с заданными параметрами: секундами и наносекундами,
	// прошедшими с начала эпохи Unix — 01.01.1970 г.
	unixTime := time.Unix(
		150000, // секунды
		1,      // наносекунды
	)
	fmt.Println(unixTime, "- unix time")

	fmt.Println()

	// парсит дату и время в строковом представлении
	firstTime, _ := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
	fmt.Println(firstTime, "- parsed time")

	// LoadLocation находит временную зону в справочнике IANA
	loc, _ := time.LoadLocation("Asia/Yekaterinburg")
	fmt.Println(loc)

	// парсит дату и время в строковом представлении с отдельным указанием временной зоны
	secondTime, _ := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
	fmt.Println(secondTime, "- parsed time with location")
	fmt.Println(secondTime.Format("2006-01-02 15:04:05"), "- formatted string")
}
