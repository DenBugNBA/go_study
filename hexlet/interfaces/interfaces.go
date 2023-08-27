package main

import (
	"fmt"
)

// объявление интерфейса
type Printer interface {
	Print()
}

// нигде не указано, что User реализует интерфейс Printer
type User struct {
	email string
}

// структура User имеет метод Print, как в интерфейсе Printer
// Следовательно, во время компиляции запишется связь между User и Printer
func (u *User) Print() {
	fmt.Println("My email is", u.email)
}

// функция принимает как аргумент интерфейс Printer
func TestPrint(p Printer) {
	p.Print()
}

func main() {
	// в функцию TestPrint передается структура User,
	// и так как она реализует интерфейс Printer, все работает без ошибок
	TestPrint(&User{email: "test@test.com"})

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic
	// f = i.(float64)

}
