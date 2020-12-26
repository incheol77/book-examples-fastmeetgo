package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data int32 = 0

	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			atomic.AddInt32(&data, 1)
		}(i)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			atomic.AddInt32(&data, -1)
		}(i)
	}

	wg.Wait()
	fmt.Println(data)
}
