package main

import "fmt"

func main() {
	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "Go"
	var s4 string = "Hello"

	fmt.Println(s1 == s2)
	fmt.Println(s1 + s2 + s3)
	fmt.Println("안녕하세요" + s1)
	fmt.Printf("%c\n", s4[0])
}

/*
go run addstring.go
true
한글한글Go
안녕하세요한글
H
*/
