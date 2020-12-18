package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b
	}(2, 3)

	fmt.Println(add)
}
