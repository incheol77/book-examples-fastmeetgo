package main

import "fmt"

func main() {
	var numPtr *int = new(int)
	*numPtr = 2

	fmt.Println(numPtr, &numPtr, *numPtr)
}
