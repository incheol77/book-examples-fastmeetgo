package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	l.PushBack(12.34)
	l.PushBack("abcd")
	l.PushBack(rune('한'))

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println(l.Front().Value, l.Back().Value)
}
