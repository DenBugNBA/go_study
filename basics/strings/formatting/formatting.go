package main

import "fmt"

func GenerateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}

type Person struct {
	Name string
	Age  int
}

func main() {
	// подставляем строку
	fmt.Println(fmt.Sprintf("hello %s", "Andy")) // "hello Andy"

	// число
	fmt.Println(fmt.Sprintf("there are %d kittens", 10)) // "there are 10 kittens"

	// логический тип
	fmt.Println(fmt.Sprintf("your story is %t", true)) // "your story is true"

	fmt.Println(fmt.Sprintf("rounded float: %.2f", 2.955))

	fmt.Println()

	p := &Person{Name: "Andy", Age: 18}

	// вывод значений структуры
	fmt.Println("simple struct:", p)

	// вывод названий полей и их значений
	fmt.Printf("detailed struct: %+v\n", p)

	// вывод названий полей и их значений в виде инициализации
	fmt.Printf("Golang struct: %#v\n", p)

	m := map[int]string{1: "foo", 2: "bar"}
	fmt.Printf("Type: %T\n", m) // Type: map[int]string
}
