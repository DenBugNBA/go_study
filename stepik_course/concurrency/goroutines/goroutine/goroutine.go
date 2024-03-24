package main

import "fmt"

func main() {
	// после запуска myFunc в отдельной горутине (на это указывает ключевое слово go),
	// функция main продолжает выполняться и завершается, не дожидавшись завершения выполнения всех прочих горутин,
	// соответственно myFunc просто не успевает завершить выполнение.
	go myFunc()
}

func myFunc() {
	fmt.Println("hello")
}
