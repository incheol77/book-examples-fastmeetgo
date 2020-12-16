package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16
	var num3 uint32 = math.MaxUint32
	var num4 uint64 = math.MaxUint64

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
}

/*
go run maxint.go
255
65535
4294967295
18446744073709551615
*/
