package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(width int, height int) *Rectangle {
	return &Rectangle{width, height}
}

func (rect *Rectangle) Area() int {
	return rect.width * rect.height
}

func (rect *Rectangle) Scale(fact int) {
	rect.width = rect.width * fact
	rect.height = rect.height * fact
}

func main() {
	r := NewRectangle(2, 3)
	fmt.Println(r.width, r.height)

	fmt.Println(r.Area())
	r.Scale(2)
	fmt.Println(r.width, r.height)
}
