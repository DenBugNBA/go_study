package main

import "fmt"

func changeSlice(arr []string) {
	arr[0] = "goodbye"
}

func appendSomeData(arr []string) {
	arr = append(arr, "!")
}

func main() {
	someSlice := []string{"hello", "world"}

	changeSlice(someSlice)
	fmt.Println(someSlice) // [goodbye world]

	appendSomeData(someSlice)
	fmt.Println(someSlice) // [goodbye world]
}
