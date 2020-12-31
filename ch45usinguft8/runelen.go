package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var r1 rune = '한'
	fmt.Println(utf8.RuneLen(r1))

	var s1 string = "한"
	fmt.Println(len(s1))

	var s2 string = "안녕하세요"
	fmt.Println(utf8.RuneCountInString(s2))
}
