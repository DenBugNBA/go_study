package main

import (
	"fmt"
)

func SumWorker(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		sum := 0

		for _, num := range nums {
			sum += num
		}

		sumCh <- sum
	}
}

func SumWorkerTeacher(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		sumCh <- sumNums(nums)
	}
}

func sumNums(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	return s
}

func main() {
	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)

	numsCh <- nil
	fmt.Println(<-sumCh) // 0

	numsCh <- []int{}
	fmt.Println(<-sumCh) // 0

	numsCh <- []int{10, 10, 10}
	fmt.Println(<-sumCh) // 3

	numsCh <- []int{500, 5, 10, 25}
	fmt.Println(<-sumCh) // 540

	fmt.Println("End")
}
