package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Index("Hello world", "He"))
	fmt.Println(strings.Index("Hello world", "lo"))

	fmt.Println(strings.IndexAny("Hello world", "e o"))

	var b byte = 'b'
	fmt.Println(strings.IndexByte("Hello world bbb", b))
	b = 'c'
	fmt.Println(strings.IndexByte("Hello world bbb", b))

	var r rune = '언'
	fmt.Println(strings.IndexRune("고 언어", r))

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	fmt.Println(strings.IndexFunc("go 언어", f))
	fmt.Println(strings.IndexFunc("Go language", f))
}
