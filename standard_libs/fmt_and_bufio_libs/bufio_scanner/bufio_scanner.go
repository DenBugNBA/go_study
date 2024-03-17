package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.ReadDir("."))
	file, err := os.Open("./standard_libs/fmt_and_bufio_libs/bufio_scanner/text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() { // возвращает true, пока файл не будет прочитан до конца
		fmt.Printf("%s\n", s.Text()) // s.Text() содержит данные, считанные на данной итерации
	}
}
