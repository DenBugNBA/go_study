package main

import "fmt"

func modifySliceAndAppend(nums []int) {
	nums[0] = 10           // элемент будет и в исходном слайсе
	nums = append(nums, 6) // элемент не добавится в исходный слайс
}

func appendAndModifySlice(nums []int) {
	nums = append(nums, 4)
	nums[0] = 2
	nums[1] = 1
	nums[2] = 0
}

func Remove(nums []int, i int) []int {
	if i < 0 || i >= len(nums) {
		return nums
	}

	var result []int

	result = append(result, nums[:i]...)
	result = append(result, nums[i+1:]...)

	return append(result)
}

func RemoveTeacherNoOrder(nums []int, i int) []int {
	if i < 0 || i > len(nums)-1 {
		return nums
	}

	nums[i] = nums[len(nums)-1]

	return nums[:len(nums)-1]
}

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}

	if maxLen > len(src) {
		maxLen = len(src)
	}

	cp := make([]int, maxLen)
	copy(cp, src)

	return cp
}

func main() {
	nums := make([]int, 0, 5)

	// panic: runtime error: index out of range [1] with length 0
	// nums[1] = 1

	nums = append(nums, 1)
	fmt.Println(nums) // [1]
	/*
		Передача слайса как аргумента функции происходит хитро.
		Длина и вместимость передаются по значению, но массив значений передается по ссылке.
		Вследствие этого получается неявное поведение:
		добавленные элементы не сохранятся в исходный слайс, но изменение существующих останется:
	*/

	modifySliceAndAppend(nums)
	fmt.Println(nums) // [10]

	a := []int{1, 2, 3}
	appendAndModifySlice(a)
	fmt.Println(a) // [1 2 3]

	RemoveTeacherNoOrder(a, 2)
	fmt.Println(a) // [1 2 3]

	a = Remove(a, 2)
	fmt.Println(a) // [1 2]

	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums1[5:]) // []

	fmt.Printf("len: %d\n", len(nums1)) // 5
	fmt.Printf("cap: %d\n", cap(nums1)) // 5

	// panic: runtime error: slice bounds out of range [6:5]
	//fmt.Println(nums1[6:]) // []

	// panic: runtime error: slice bounds out of range [:6] with capacity 5
	// fmt.Println(nums1[:6])

	fmt.Println(nums1[:5]) // [1 2 3 4 5]

	fmt.Println()

	for i, num := range nums1 {
		fmt.Println(i, num)
	}

	fmt.Println()

	for i := range nums1 {
		fmt.Println(i)
	}

	fmt.Println()

	var emptySlice []int
	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))
	if emptySlice == nil {
		fmt.Println("nil!")
	}

	fmt.Println()

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = s[:]
	printSlice(s)

	fmt.Println()

	b := make([]int, 0, 5) // len=0 cap=5 []

	printSlice(b)

	b = b[:] // len=0 cap=5 []

	printSlice(b)

	b = b[:cap(b)] // len=5 cap=5 [0 0 0 0 0]
	printSlice(b)

	b = b[1:] // len=4 cap=4 [0 0 0 0]

	printSlice(b)

	numsArr := [5]int{1, 2, 3, 4, 5}
	first := numsArr[:3]
	printSlice(first)
	first = numsArr[2:]
	printSlice(first)

	fmt.Println()

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// A slice does not store any data, it just describes a section of an underlying array.
	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	fmt.Println()

	var arr []string
	fmt.Println(arr) // []

	arr = nil
	fmt.Println(arr) // []

	arr = append(arr, "ddd")
	fmt.Println(arr) // [ddd]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
