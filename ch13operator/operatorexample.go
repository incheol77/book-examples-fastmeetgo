package main

import (
	"fmt"
)

func main() {
	var num1 uint32 = 10
	var num2 float32 = 10.2
	var num3 complex64 = 2 + 3i

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	num1++
	num2++
	num3++

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	num1--
	num2--
	num3--

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	/*
		10.2
		(2+3i)
		11
		11.2
		(3+3i)
		10
		10.2
		(2+3i)
	*/
}
