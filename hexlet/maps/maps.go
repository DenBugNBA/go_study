package main

import "fmt"

func UniqueUserIDs(userIDs []int64) []int64 {
	// пустая структура struct{} — это тип данных, который занимает 0 байт
	// используется, когда нужно проверять в мапе только наличие ключа
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, uid := range userIDs {
		_, exists := processed[uid]
		if !exists {
			uniqUserIDs = append(uniqUserIDs, uid)
			processed[uid] = struct{}{}
		}
	}

	return uniqUserIDs
}

func MostPopularWord(words []string) string {
	wordTocount := make(map[string]int)
	currentMaxCount := 0

	for _, word := range words {
		wordTocount[word]++

		if wordTocount[word] > currentMaxCount {
			currentMaxCount = wordTocount[word]
		}
	}

	for _, word := range words {
		if wordTocount[word] == currentMaxCount {
			return word
		}
	}

	return ""
}

func main() {
	engToRusEmpty := make(map[string]string)
	fmt.Println(engToRusEmpty) // map[]

	// no-op
	delete(engToRusEmpty, "word")

	rus := engToRusEmpty["word"]
	fmt.Println(rus) //

	rus, ok := engToRusEmpty["word"]
	fmt.Println(rus, ok) //  false

	fmt.Println()

	engToRus := make(map[string]string, 10)
	fmt.Println(engToRus)      // map[]
	fmt.Println(len(engToRus)) // 0

	m1 := map[int]string{}
	fmt.Println(m1) // map[]

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)          // map[bar:2 foo:1]
	fmt.Println(m2["hello"]) // 0

	idToName := map[int64]string{1: "Alex", 2: "Dan", 3: "George"}

	// первый аргумент — ключ, второй — значение
	for id, name := range idToName {
		fmt.Println("id: ", id, "name: ", name)
	}

	for id := range idToName {
		fmt.Println("id: ", id)
	}

	fmt.Println()

	var i int
	fmt.Println(i) // 0
	i = 1
	fmt.Println(i) // 1

	var words []string
	words = append(words, "hello")
	fmt.Println(words) // [hello]
}
