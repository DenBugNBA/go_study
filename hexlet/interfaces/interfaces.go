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

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

type Human interface {
	SayHello()
}

func main() {
	// в функцию TestPrint передается структура User,
	// и так как она реализует интерфейс Printer, все работает без ошибок
	TestPrint(&User{email: "test@test.com"})

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// type assertion
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic
	// f = i.(float64)

	fmt.Println()

	// переменная типа Human, которая будет ссылаться на объект типа Person
	var h Human = Person{Name: "John", Age: 25}
	h.SayHello()

	fmt.Println()

	// может хранить значения любых типов данных
	var values []interface{}
	values = append(values, "hello")
	values = append(values, 42)
	values = append(values, true)
	fmt.Println(values)
}

func anyType(val any) {}
