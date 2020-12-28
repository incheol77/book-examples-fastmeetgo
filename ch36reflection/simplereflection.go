package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	var num int64 = 1
	fmt.Println(reflect.TypeOf(num))

	var s string = "hello world"
	fmt.Println(reflect.TypeOf(s))

	f := 12.34
	fmt.Println(reflect.TypeOf(f))

	var d Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(d))
}
