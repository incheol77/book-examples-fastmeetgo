package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data = []int{}
	count := 1000

	go func() {
		for i := 0; i < count; i++ {
			data = append(data, i)
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < count; i++ {
			data = append(data, i)
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("len data : ", len(data))
}
