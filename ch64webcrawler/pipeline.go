package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(n int, done <-chan struct{}, jobs <-chan int, c chan<- string) {
	for j := range jobs {
		select {
		case c <- fmt.Sprintf("Worker: %d, Jobs: %d", n, j):
		case <-done:
			return
		}
	}
}

func main() {
	jobs := make(chan int)
	done := make(chan struct{})
	c := make(chan string)

	var wg sync.WaitGroup
	const numWorkers = 5
	wg.Add(numWorkers)

	// create workers
	for i := 0; i < numWorkers; i++ {
		go func(n int) {
			worker(n, done, jobs, c)
			wg.Done()
		}(i)
	}

	// wait for workers
	go func() {
		wg.Wait()
		close(c)
	}()

	// clear channels
	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i
			time.Sleep(time.Millisecond)
		}
		close(done)
		close(jobs)
	}()

	for r := range c {
		fmt.Println(r)
	}
}
