package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s []byte = make([]byte, 0)

	s = strconv.AppendBool(s, true)
	fmt.Println(s, string(s))

	s = strconv.AppendFloat(s, 1.23, 'f', -1, 32)
	fmt.Println(s, string(s))

	s = strconv.AppendInt(s, 123, 10)
	fmt.Println(s, string(s))

	var err error

	var b1 bool
	b1, err = strconv.ParseBool("true")
	fmt.Println(b1, err)

	var num1 int64
	num1, err = strconv.ParseInt("12345", 10, 32)
	fmt.Println(num1, err)
}
