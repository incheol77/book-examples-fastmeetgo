package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(i int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(i)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	fmt.Scanln()
}