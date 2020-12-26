package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan int, 4)
	count := 5

	go func() {
		for i := 0; i < count; i++ {
			c <- i
			fmt.Println("go routine")
		}
		close(c)
	}()

	for i := range c {
		fmt.Println("main routine : ", i)
	}
}
