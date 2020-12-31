package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("He", "Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[a-z0-9A-Z]+", "Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[a-z0-9A-Z]+", "")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[a-z0-9A-Z]*", "")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[0-9]+-[0-9]+", "1-5")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[0-9]+\\+[0-9]*", "1+")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[A-Z]+", "hello")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[^A-Z]+", "hello")
	fmt.Println(matched)
}
