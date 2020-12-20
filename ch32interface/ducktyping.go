package main

import "fmt"

type Duck struct {
}

func (d Duck) Quack() {
	fmt.Println("Duck quack")
}

func (d Duck) Feather() {
	fmt.Println("Duck feather")
}

type Person struct {
}

func (p Person) Quack() {
	fmt.Println("Person quack")
}

func (p Person) Feather() {
	fmt.Println("Person feather")
}

type Quacker interface {
	Quack()
	Feather()
}

func CreatureInForest(q Quacker) {
	q.Quack()
	q.Feather()
}

func main() {
	d := Duck{}
	p := Person{}
	CreatureInForest(d)
	CreatureInForest(p)

	if v, ok := interface{}(d).(Quacker); ok {
		fmt.Println(v, true)
	}
}
