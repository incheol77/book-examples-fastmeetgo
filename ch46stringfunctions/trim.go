package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Trim("~~~Hello~~, world~~", "~~"))
	fmt.Println(strings.TrimLeft("~~~Hello~~, world~~", "~~"))
	fmt.Println(strings.TrimRight("~~~Hello~~, world~~", "~~"))

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	var s string = "안녕 go 언어 language 만세"
	fmt.Println(strings.TrimFunc(s, f))
	fmt.Println(strings.TrimLeftFunc(s, f))
	fmt.Println(strings.TrimRightFunc(s, f))

	fmt.Println(strings.TrimSuffix("Hello world", "ld"))

	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("<div><span>Hello, world</span></div>"))
}
