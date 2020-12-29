package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			fmt.Println("producer : Lock")
			c <- true
			fmt.Println("wait begin : ", n)
			cond.Wait()
			fmt.Println("wait end : ", n)
			mutex.Unlock()
			fmt.Println("producer : UnLock")
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("consumer : Lock")
		fmt.Println("signal : ", i)
		cond.Signal()
		mutex.Unlock()
		fmt.Println("consumer : UnLock")
	}

	fmt.Scanln()
}