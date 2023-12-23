package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	// 27 >= 20
	// 9 20

	fmt.Println()

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	// Linux.

	fmt.Println()

	x := 10

	switch x {
	// default всегда выполняется последним независимо от расположения в конструкции
	default:
		fmt.Println("default case")
	case 10:
		fmt.Println("case 10")
	} // case 10

	fmt.Println()

	i := 1

	switch i {
	case 0:
	case func() int { fmt.Println("Hi"); return 1 }():
		fmt.Println("Success")
	}
	// Hi
	// Success

	fmt.Println()

	// выражение отсутствует. Для компилятора выглядит как: switch true
	switch {
	default:
		fmt.Println("default case")
	case x == 10:
		fmt.Println("equal 10 case")
		fallthrough
	case x <= 10:
		fmt.Println("less or equal 10 case")
	}
	// equal 10 case
	// less or equal 10 case

	fmt.Println()

	x = 5

	switch {
	case x == 5:
		fmt.Println("Bob")
		fallthrough
	case x <= 5:
		fmt.Println("Alice")
		fallthrough
	case x > 55:
		fmt.Println("Joe")
		fallthrough
	default:
		fmt.Println("Default")
	}
	// Bob
	// Alice
	// Joe
	// Default

	fmt.Println()

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println()

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	do(21)
	do("hello")
	do(true)

	fmt.Println()

	num := 3
	word := "korov"

	switch lastNum := num % 10; lastNum {
	case 1:
		word += "a"
	case 2, 3, 4:
		word += "y"
	}

	fmt.Printf("%d %s", num, word)

}
