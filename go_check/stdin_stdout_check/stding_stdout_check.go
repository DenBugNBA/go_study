package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in)

	var n int
	fmt.Fscan(in, &n)

	fmt.Println(n, "- n")

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
