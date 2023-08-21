package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSum([]int{1, 2, 3}, []int{10, 20, 50})) // [10 20 50]
}

// суммирует значения каждого слайса nums и возвращает тот, который имеет наибольшую сумму
func maxSum(nums1, nums2 []int) []int {
	// канал для результата первой суммы
	s1Ch := make(chan int)
	go sumParallel(nums1, s1Ch)

	// канал для результата второй суммы
	s2Ch := make(chan int)
	go sumParallel(nums2, s2Ch)

	// присваиваем результаты в переменные. Здесь программа будет заблокирована, пока не придут результаты из обоих каналов.
	s1, s2 := <-s1Ch, <-s2Ch

	if s1 > s2 {
		return nums1
	}

	return nums2
}

func sumParallel(nums []int, resCh chan int) {
	s := 0
	for _, n := range nums {
		s += n
	}

	// результат суммы передаем в канал
	resCh <- s
}
