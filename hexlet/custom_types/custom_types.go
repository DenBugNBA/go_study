package main

import "fmt"

type NumCount int
type errorCode string
type counter int

// передается указатель, чтобы можно было изменить состояние счетчика "c"
func (c *counter) inc() {
	*c++
}

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	popularity := make(map[uint8]int)
	for _, p := range pl {
		popularity[p.Age]++
	}

	return popularity
}

func main() {
	nc := NumCount(len([]int{1, 2, 3}))
	fmt.Println(nc) // 3

	ec := errorCode("internal")
	fmt.Println(ec)         // internal
	fmt.Println(string(ec)) // internal

	c := counter(0)
	(&c).inc() // передается указатель на счетчик &c, так как функция "inc()" работает с указателями
	(&c).inc()

	fmt.Println(c) // 2

	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}

	fmt.Println(pl.GetAgePopularity()) // map[18:2 44:1]
}
