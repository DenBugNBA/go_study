package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for temp := l.Front(); temp != nil; temp = temp.Next() {
		fmt.Printf("%v ", temp.Value)
	}
	fmt.Println()
}

func main() {
	// Создаем список
	myList := list.New()
	myList.PushBack(0)
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	// получаем указатель на элемент который добавили (*Element)
	lastElem := myList.PushBack(4)
	printList(myList)

	// удаляем элемент '3' по указателю (*Element)
	myList.Remove(lastElem)
	printList(myList)

	// так же можем воспользоваться методом Front() / Back() чтобы получить первый/последний элемент
	myList.Remove(myList.Front())

	fmt.Printf("\nПосле удаления Front():\n")
	printList(myList)
}
