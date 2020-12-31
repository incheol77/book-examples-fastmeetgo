package main

import (
	"fmt"
	"regexp"
)

func main() {
	var validEmail, _ = regexp.Compile(
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
	)

	fmt.Println(validEmail.MatchString("hello@email.com"))
	fmt.Println(validEmail.MatchString("HELLO@email.com"))
	fmt.Println(validEmail.MatchString("hello-world+@email.com"))

	fmt.Println(validEmail.MatchString("hello-world+@email.coooom"))
	fmt.Println(validEmail.MatchString("hello-world+@email"))
}
