package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	var f float64 = 1.234
	var d Data = Data{1, 2}

	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	fmt.Println(t)
	fmt.Println(v)

	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.Kind())

	fmt.Println(v.Type())
	fmt.Println(v.Float())

	td := reflect.TypeOf(d)
	vd := reflect.ValueOf(d)
	fmt.Println(td.Name(), td.Kind())
	fmt.Println(vd)
}
