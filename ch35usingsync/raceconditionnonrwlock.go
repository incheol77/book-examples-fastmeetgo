package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data int = 0
	count := 3

	// write goroutine
	go func() {
		for i := 0; i < count; i++ {
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// read goroutine
	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("read 1 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("read 2 : ", data)
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
