package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	el1 := List[string]{val: "hello"}
	el2 := List[string]{next: &el1, val: "world"}

	fmt.Println(el2.next.val)

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	fmt.Printf("Generic Sum with Constraint: %v\n",
		SumNumbers(ints))
}
