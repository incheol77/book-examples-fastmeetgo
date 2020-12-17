package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println(s1, " : size = ", len(s1))
	fmt.Println(s1, " : character size = ", utf8.RuneCountInString(s1))
	fmt.Println(s2, " : size = ", len(s2))
}

/*
go run stringlength.go
한글  : size =  6
한글  : character size =  2
Hello  : size =  5
*/
