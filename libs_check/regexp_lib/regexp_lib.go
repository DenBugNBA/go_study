package main

import (
	"fmt"
	"regexp"
)

func ContainsOnlyDigits(str string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(str)
}

func main() {
	fmt.Println(ContainsOnlyDigits("123"))  // true
	fmt.Println(ContainsOnlyDigits("123?")) // false
}
