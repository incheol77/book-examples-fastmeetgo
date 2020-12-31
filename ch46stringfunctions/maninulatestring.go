package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := []string{"Hello", "World"}
	fmt.Println(strings.Join(s1, " + "))

	s2 := strings.Split("Hello, world", ",")
	fmt.Println(s2[0], s2[1])

	s3 := strings.Fields("Hello, world")
	fmt.Println(s3[0])

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	fmt.Println(strings.FieldsFunc("hangul한글korea", f))

	fmt.Println(strings.Repeat("repeat", 2))

	fmt.Println(strings.Replace("Hello world", "world", "Go", 1))
	fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 1))
	fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 2))
}
