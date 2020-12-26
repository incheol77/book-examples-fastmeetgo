package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	var data int = 0

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			data++
		}(i)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			data--
		}(i)
	}

	wg.Wait()
	fmt.Println(data)
}
