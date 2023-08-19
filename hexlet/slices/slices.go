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
	fmt.Println(emptySlice)
}
