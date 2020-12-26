package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("go routine : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("main routine : ", i)
	}
}