package main

import "fmt"

func increase(numPtr *int) int {
	*numPtr++

	return *numPtr
}

func main() {
	num := 1
	fmt.Println(num)
	increase(&num)
	fmt.Println(num)
}
