package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var str string
	s := "1 2 hello"
	n, _ := fmt.Sscanln(s, &num1, &num2, &str)
	fmt.Println(n, num1, num2, str)
}
