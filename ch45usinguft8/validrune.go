package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var v1 []byte = []byte("안녕하세요")
	fmt.Printf("%s : ", v1)
	fmt.Println(utf8.Valid(v1))

	v2 := []byte{0xff, 0xc1, 0xd1}
	fmt.Printf("%s : ", v2)
	fmt.Println(utf8.Valid(v2))

	var v3 rune = '한'
	fmt.Println(utf8.ValidRune(v3))
}
