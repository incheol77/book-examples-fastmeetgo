package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	var rects []*Rectangle = make([]*Rectangle, 3)

	rects[0] = new(Rectangle)
	rects[1] = &Rectangle{2, 5}
	rects[2] = &Rectangle{width: 3, height: 5}

	for _, r := range rects {
		fmt.Println(r.width, r.height)
	}
}
