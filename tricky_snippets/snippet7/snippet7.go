package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(X2())
}

func X1() (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = e.(error)
		}
	}()
	panic(errors.New("my error"))
}

func X2() (err error) {
	defer recoverSomething(&err)

	panic(errors.New("my error"))
}

func recoverSomething(err *error) {
	e := recover()

	if e != nil {
		errRecovered := e.(error)
		errPointer := &errRecovered
		err = errPointer
	}

	fmt.Println(err)
	fmt.Println(*err)
}
