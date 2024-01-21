package main

import (
	"fmt"
	"math"
)

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	fmt.Println(MinInt(5, 4)) // 4
}
