package main

import "fmt"

func main() {
	var num1 int = 1
	var num2 float32 = 2.2

	fmt.Println("a + b = ", (float32)(num1)+num2)
	fmt.Println("a + b = ", num1+(int)(num2))
}

/*
go run typecasting01.go
a + b =  3.2
a + b =  3
*/
