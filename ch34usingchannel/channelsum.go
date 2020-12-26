package main

import "fmt"

func num(a, b int) <-chan int {
	in := make(chan int)
	go func() {
		in <- a
		in <- b
		close(in)
	}()
	return in
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for n := range in {
			r = r + n
		}
		out <- r
	}()
	return out
}

func main() {
	in := num(1, 2)
	out := sum(in)
	fmt.Println(<-out)
}
