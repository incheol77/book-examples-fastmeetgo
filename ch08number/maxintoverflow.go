package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 uint8 = math.MaxUint8 + 1
	var num2 uint16 = math.MaxUint16 + 1
	var num3 uint32 = math.MaxUint32 + 1
	var num4 uint64 = math.MaxUint64 + 1

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
}

/*
go run maxintoverflow.go
# command-line-arguments
./maxintoverflow.go:9:33: constant 256 overflows uint8
./maxintoverflow.go:10:35: constant 65536 overflows uint16
./maxintoverflow.go:11:35: constant 4294967296 overflows uint32
./maxintoverflow.go:12:35: constant 18446744073709551616 overflows uint64
*/
