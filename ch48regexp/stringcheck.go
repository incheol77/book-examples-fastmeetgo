package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("^Hello", "Hello world")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("^hello", "Hello world")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("world$", "Hello world")
	fmt.Println(matched)
}
