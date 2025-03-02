package main

import "fmt"

func main() {
	nums := make([]int, 0, 5)
	_ = append(nums, 1)

	fmt.Println(nums)     // []
	fmt.Println(nums[:5]) // [1 0 0 0 0]
}
