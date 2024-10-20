package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Создаем контейнер list и добавляем в него элементы
	myList := list.New()
	myList.PushBack(2)
	myList.PushBack(15)
	myList.PushBack(3)
	myList.PushBack(9)
	myList.PushBack(12)

	// Проходимся по элементам и удаляем те, которые меньше 10
	for e := myList.Front(); e != nil; {
		next := e.Next() // Запоминаем следующий элемент перед удалением текущего
		if e.Value.(int) < 10 {
			myList.Remove(e) // Удаляем текущий элемент из списка
		}
		e = next // Переходим к следующему элементу
	}

	// Выводим список после удаления элементов
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
