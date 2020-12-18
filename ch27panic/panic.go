package main

import "fmt"

func f() {
	defer func() {
		// r := recover()
		r := "no recover"
		fmt.Println(r)
	}()
	panic("panic!!")
}

func main() {
	f()
	fmt.Println("Hello, world")
}
