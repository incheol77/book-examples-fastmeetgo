package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	grade int
	score float64
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByGrade struct {
	Students
}

func (s ByGrade) Less(i, j int) bool {
	return s.Students[i].grade < s.Students[j].grade
}

type ByScore struct {
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score
}

func main() {
	s := Students{
		{"Maria", 2, 83.2},
		{"Andrew", 4, 92.3},
		{"John", 1, 75.7},
	}

	fmt.Println(s)

	fmt.Println("---- By name ----")
	sort.Sort(s)
	fmt.Println(s)
	sort.Sort(sort.Reverse(s))
	fmt.Println(s)

	fmt.Println("---- By grade ----")
	sort.Sort(ByGrade{s})
	fmt.Println(s)
	sort.Sort(sort.Reverse(ByGrade{s}))
	fmt.Println(s)

	fmt.Println("---- By score ----")
	sort.Sort(ByScore{s})
	fmt.Println(s)
	sort.Sort(sort.Reverse(ByGrade{s}))
	fmt.Println(s)
}
