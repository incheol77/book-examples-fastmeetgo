package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello world", "wo"))
	fmt.Println(strings.Contains("Hello world", "w o"))
	fmt.Println(strings.ContainsAny("Hello world", "w o"))

	fmt.Println(strings.Count("Hello world", "l"))

	var r rune = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r))

	fmt.Println(strings.HasPrefix("Hello world", "He"))
	fmt.Println(strings.HasPrefix("Hello world", "he"))

	fmt.Println(strings.HasSuffix("Hello world", "ld"))
}
