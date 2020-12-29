package main

import (
	calc "books-examples-go/book-examples-fastmeetgo/ch39makepackage/calc"
	"fmt"
)

func main() {
	add := calc.Sum(1, 2)
	fmt.Println(add)
}
