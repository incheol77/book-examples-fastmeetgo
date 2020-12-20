package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) hello() {
	fmt.Println("Person : ", p.name, p.age)
}

type Student struct {
	p      Person
	school string
	grade  int
}

func (s *Student) hello() {
	s.p.hello()
	fmt.Println("Student : ", s.school, s.grade)
}

func main() {
	p := Person{"Tom", 30}
	s := Student{Person{"Jane", 40}, "MIT", 3}

	p.hello()
	s.hello()
}
