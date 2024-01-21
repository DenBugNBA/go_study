package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Быть или не быть.
	text, _ := in.ReadString('\n')

	fmt.Println(text) // Быть или не быть.
}
