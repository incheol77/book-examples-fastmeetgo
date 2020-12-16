package main

import "fmt"

func main() {
	var num1 complex64 = 1 + 2i
	var num2 complex64 = 4.321 + 5.432i

	fmt.Println(num1)
	fmt.Println(num2)

	var num3 complex64 = 1.e+3i
	var num4 complex64 = 7.276443 - 10i

	fmt.Println(num3)
	fmt.Println(num4)

	var num5 complex128 = 1 + 2i
	var num6 complex128 = 5.3244e-10 + .12345E+2i

	fmt.Println(num5)
	fmt.Println(num6)

	var r1 float32 = real(num2)
	var i1 float32 = imag(num2)

	fmt.Println(r1)
	fmt.Println(i1)
}

/*
go run complexnumber.go
(1+2i)
(4.321+5.432i)
(0+1000i)
(7.276443-10i)
(1+2i)
(5.3244e-10+12.345i)
4.321
5.432
*/
