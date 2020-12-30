package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var str string

	file1, _ := os.Open("hello.txt")
	defer file1.Close()
	fmt.Fscan(file1, &num1, &num2, &str)
	fmt.Println(num1, num2, str)
}
