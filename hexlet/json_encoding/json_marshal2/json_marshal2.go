package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	val := map[string]int{"one": 1, "two": 2, "three": 3, "четыре": 4}
	bytes, _ := json.Marshal(val)

	fmt.Println(bytes)
	fmt.Println(string(bytes)) // {"one":1,"three":3,"two":2,"четыре":4}

	arr := []string{"one", "two", "три"}
	jsonBytes, _ := json.Marshal(arr)

	fmt.Println(jsonBytes)
	fmt.Println(string(jsonBytes)) // ["one","two","три"]
}
