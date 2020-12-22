package main

import "fmt"

func hello() {
	fmt.Println("go routine : hello")
}

func main() {
	go hello()

	fmt.Println("main : hello")
	fmt.Scanln()
}
