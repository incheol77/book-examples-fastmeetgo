package main

import "fmt"

func f() {
	arr := [...]int{1, 2}
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	f()
	fmt.Println("Hello, world")
}
