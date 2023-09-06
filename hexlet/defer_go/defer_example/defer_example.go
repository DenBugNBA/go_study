package main

import "fmt"

func F(x int) {
	fmt.Println(x)
}

func main() {
	i := 1

	defer F(i)
	defer func() { F(i) }()
	defer func(i int) { F(i) }(i)

	i = 2
}

/*
1
2
1
*/
