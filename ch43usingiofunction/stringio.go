package main

import (
	"fmt"
)

func main() {
	var s1 string
	s1 = fmt.Sprint(1, 1.2, "hello world")
	fmt.Print(s1)
}
