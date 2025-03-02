package main

import "fmt"

func main() {
	nums := make([]int, 1, 3)
	fmt.Println(nums) // [0] len=1 cap=3

	appendSlice(nums, 1)
	fmt.Println(nums) // [0] len=1 cap=3

	fmt.Println(nums[:3]) // [0 1 0] - 1 есть в исходном массиве

	copySlice(nums, []int{2, 3})
	fmt.Println(nums) // [2] len=1 cap=3

	fmt.Println(nums[:3]) // [2 1 0]

	// mutateSlice(nums, 1, 4) // panic: index out of range [1] with length 1
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func copySlice(sl []int, cp []int) {
	copy(sl, cp)
	fmt.Println("in copySlice sl: ", sl) // [2]
	fmt.Println("in copySlice cp: ", cp) // [2 3]
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
