package main

import "fmt"

// можно не указывать значение следующей константы по порядку
// (значение будет скопировано):
const (
	A int = 45
	B
	C float32 = 3.3
	D
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(A, B, C, D) // 45 45 3.3 3.3

	const pi float64 = 3.1415
	fmt.Println(pi) // 3.1415

	fmt.Println(Sunday) // 0

	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	fmt.Println(c0, c1, c2) // 0 1 2

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // пропускаем 7
		Add
	)
	fmt.Println(Sunday, Saturday, Add) // 0 6 8

	const (
		u         = iota * 42 // u == 0 (индекс - 0, поэтому 0 * 42 = 0)
		v float64 = iota * 42 // v == 42.0 (индекс - 1, поэтому 1.0 * 42 = 42.0)
		w         = iota * 42 // w == 84  (индекс - 2, поэтому 2 * 42 = 84)
	)
	fmt.Println(u, v, w) // 0 42 84

	// переменные ни в одном блоке const, поэтому индекс не увеличился
	const x = iota
	const y = iota
	fmt.Println(x, y) // 0 0
}
