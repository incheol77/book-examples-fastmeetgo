package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string
	s1 = strconv.FormatBool(true)
	fmt.Println(s1)

	var s2 string
	s2 = strconv.FormatFloat(1.3, 'f', -1, 32)
	fmt.Println(s2)

	var s3 string
	s3 = strconv.FormatInt(-10, 10)
	fmt.Println(s3)

	var s4 string
	s4 = strconv.FormatUint(32, 16)
	fmt.Println(s4)
}
