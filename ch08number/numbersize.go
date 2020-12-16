package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 uint8 = 1
	var num2 uint16 = 1
	var num3 uint32 = 1
	var num4 uint64 = 1

	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))
}

/*
go run numbersize.go
1
2
4
8
*/
