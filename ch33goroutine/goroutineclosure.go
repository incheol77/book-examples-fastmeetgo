package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := "hello, world"

	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println("go routine : ", s, n)
		}(i)
	}

	fmt.Println("main routine")
	fmt.Scanln()
}