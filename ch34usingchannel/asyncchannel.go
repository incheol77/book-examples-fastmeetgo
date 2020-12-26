package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	done := make(chan bool, 2)
	count := 8

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("go routine : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("main routine : ", i)
	}
}