package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow, "- time now")

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
}
