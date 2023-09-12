package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"  hello   ", "   world "}
	TrimAllWithoutCycle(strs)
	fmt.Println(strs)
}

func TrimAll(strs []string) {
	for i, s := range strs {
		strs[i] = strings.TrimSpace(s)
	}
}

func TrimAllWithoutCycle(strs []string) {
	i := len(strs)
LOOP:
	if i--; i >= 0 {
		strs[i] = strings.TrimSpace(strs[i])
		goto LOOP
	}
}
