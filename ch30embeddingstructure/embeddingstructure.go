package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) hello() {
	fmt.Println("Person hello")
}

type Student struct {
	Person
	school string
	grade  int
}

func (s *Student) hello() {
	fmt.Println("Student hello")
}

func main() {
	p := Person{"Tom", 30}
	s := Student{Person{"Jane", 40}, "MIT", 3}

	p.hello()
	fmt.Println("----")
	s.hello()
}
