package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	epsilon := 10e-14
	fmt.Println(a)
	fmt.Println(math.Abs(a-9.0) <= epsilon)
}

/*
go run realnumber02.go
9.000000000000004
true
*/
