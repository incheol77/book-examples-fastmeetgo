package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data = []int{}
	var mutex = new(sync.Mutex)
	count := 1000

	go func() {
		for i := 0; i < count; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < count; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("len data : ", len(data))
}
