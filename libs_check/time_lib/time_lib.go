package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow)

	timeNowUnix := timeNow.Unix() // number of seconds elapsed since January 1, 1970, UTC.
	fmt.Println(timeNowUnix)

	const minTime = 94608000

	fmt.Println(time.Unix(minTime, 0), "- min time")

	fmt.Println()

	randomSeconds := rand.Int63n(timeNowUnix-minTime) + minTime
	fmt.Println(randomSeconds, "- random seconds")

	randomTime := time.Unix(randomSeconds, 0)
	fmt.Println(randomTime, "- random time")

	fmt.Println(randomTime.String())
}
