package main

import (
	"errors"
	"fmt"
)

func validateName(name string) error {
	if name == "" {
		// errors.New создает новый объект ошибки
		return errors.New("empty name")
	}

	if len([]rune(name)) > 50 {
		return errors.New("a name cannot be more than 50 characters")
	}

	return nil
}

type TimeoutErr struct {
	msg string
}

// структура TimeoutErr реализует интерфейс error
// и может быть использована как обычная ошибка
func (e *TimeoutErr) Error() string {
	return e.msg
}

func main() {
	err := errors.New("my error")
	fmt.Println(err)
}
