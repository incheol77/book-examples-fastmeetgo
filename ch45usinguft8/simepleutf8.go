package main

import (
	"fmt"
	"unicode"
)

func main() {
	var r1 rune = '한'
	fmt.Println(unicode.Is(unicode.Hangul, r1))
	fmt.Println(unicode.Is(unicode.Latin, r1))

	var r2 rune = 'a'
	fmt.Println(unicode.Is(unicode.Hangul, r2))
	fmt.Println(unicode.Is(unicode.Latin, r2))

	fmt.Println(unicode.In(r1, unicode.Hangul, unicode.Latin))
}
