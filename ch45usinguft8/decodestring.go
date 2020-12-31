package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello!"
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%c\n", s[len(s)-1])

	s = "안녕하세요"
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%c\n", s[len(s)-1])

	c, _ := utf8.DecodeRuneInString(s)
	fmt.Printf("%c\n", c)

	c, _ = utf8.DecodeLastRuneInString(s)
	fmt.Printf("%c\n", c)
}
