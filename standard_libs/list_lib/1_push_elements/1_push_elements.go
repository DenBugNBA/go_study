package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Создаем список
	myList := list.New()

	// Добавляем каждый элемент в конец
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	// Пробегаемся по списку и печатаем не пустые элементы
	// Мы не можем пробежаться привычным способом как с массивами,
	// поэтому придется использовать метод Front()
	// которая вернет первый элемент и затем с помощью Next
	// получать следующий элемент пока он не будет равен nil, что означает конец списка
	for temp := myList.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}

	fmt.Println()

	// Создаем список
	myList2 := list.New()

	// Добавляем каждый элемент в начало
	myList2.PushFront(1)
	myList2.PushFront(2)
	myList2.PushFront(3)

	// Пробегаемся по списку и печатаем не пустые элементы
	for temp := myList2.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
}
