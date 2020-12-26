package main

import "fmt"

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
}

func consumer(c <-chan int) {
	for n := range c {
		fmt.Println(n)
	}
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
