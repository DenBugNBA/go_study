package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(XCorrect())
}

func XOriginal() (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = e.(error)
		}
	}()
	panic(errors.New("my error"))
}

func XCorrect() (err error) {
	defer suppress(&err)
	panic(errors.New("my error"))
}

// Но так считается не совсем правильно, считается,
// что надо все же прокладывать функцией и recover вызывать
// только в defer, но это уже не вопрос задачи
func suppress(err *error) {
	e := recover()
	if e != nil {
		*err = e.(error)
	}
}
