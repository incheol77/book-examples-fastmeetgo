package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("hello.txt")
	defer file1.Close()
	fmt.Fprint(file1, 1, 1.1, "hello")

	file2, _ := os.Create("world.txt")
	defer file2.Close()
	fmt.Fprintf(file2, "%d %f %s\n", 2, 2.2, "world")
}
