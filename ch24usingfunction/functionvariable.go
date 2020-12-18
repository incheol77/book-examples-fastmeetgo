package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func main() {
	f1 := []func(int, int) int{
		sum,
		sub,
	}
	fmt.Println(f1[0](2, 1))
	fmt.Println(f1[1](2, 1))

	f2 := map[string]func(int, int) int{
		"sum": sum,
		"sub": sub,
	}
	fmt.Println(f2["sum"](2, 1))
	fmt.Println(f2["sub"](2, 1))
}
