package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int
	var err error

	num1, err = strconv.Atoi("100")
	fmt.Println(num1, err)

	num1, err = strconv.Atoi("10t")
	fmt.Println(num1, err)

	var s string
	s = strconv.Itoa(1000)
	fmt.Println(s)
}
