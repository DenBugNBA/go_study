package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es")) // true

	fmt.Println(strings.Count("test", "t")) // 2

	fmt.Println(strings.Fields("  one    two\tthree four ")) // [one two three four]

	fmt.Println(strings.HasPrefix("test", "te")) // true

	fmt.Println(strings.HasSuffix("test", "st")) // true

	fmt.Println(strings.Index("test", "e")) // 1

	fmt.Println(strings.Join([]string{"hello", "world"}, "-")) // hello-world

	fmt.Println(strings.Repeat("a", 5)) // aaaaa

	fmt.Println(strings.Replace("blanotblanot", "not", "***", -1)) // bla***bla***
	fmt.Println(strings.Replace("blanotblanot", "", "***", -1))    // ***b***l***a***n***o***t***b***l***a***n***o***t***

	fmt.Println(strings.ReplaceAll("one two three", " ", "_")) // one_two_three

	fmt.Println(strings.Split("a-b-c-d-e", "-")) // [a b c d e]

	fmt.Println(strings.ToLower("TEST")) // test

	fmt.Println(strings.ToUpper("test")) // TEST

	fmt.Println(strings.Trim("tetstet", "te")) // s

	fmt.Println(strings.TrimSpace(" hello ")) // "hello"

	sb := strings.Builder{}

	sb.WriteString("hello")
	sb.WriteString(" ")
	sb.WriteString("world")

	fmt.Println(sb.String()) // "hello world"
}
