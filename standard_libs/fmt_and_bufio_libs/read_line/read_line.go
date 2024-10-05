package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	fmt.Println(text)
}
