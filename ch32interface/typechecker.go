package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

type Student struct {
	school string
	grade  int
}

func adapter(arg interface{}) string {
	switch arg.(type) {
	case Person:
		p := arg.(Person)
		return "var Person : " + p.name + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return "ptr Person : " + p.name + strconv.Itoa(p.age)
	case Student:
		s := arg.(Student)
		return "var Student : " + s.school + strconv.Itoa(s.grade)
	case *Student:
		s := arg.(*Student)
		return "ptr Student : " + s.school + strconv.Itoa(s.grade)
	default:
		return "Error"
	}
}

func main() {
	pVar := Person{"Tom", 40}
	pPtr := &Person{"Jane", 20}
	sVar := Student{"AAA School", 3}
	sPtr := &Student{"BBB School", 2}

	fmt.Println(adapter(pVar))
	fmt.Println(adapter(pPtr))
	fmt.Println(adapter(sVar))
	fmt.Println(adapter(sPtr))
}
