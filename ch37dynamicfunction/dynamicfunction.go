package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value {
	fmt.Println("hello world in h")
	return nil
}

func main() {
	var hello func()
	fn := reflect.ValueOf(&hello).Elem()
	v := reflect.MakeFunc(fn.Type(), h)
	fn.Set(v)

	hello()
}
