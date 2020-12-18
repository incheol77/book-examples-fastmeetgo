package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var data []byte
	var err error

	data, err = ioutil.ReadFile("./data.txt")

	if err == nil {
		fmt.Println(data)
		fmt.Printf("%s\n", data)
	}

	if data2, err2 := ioutil.ReadFile("./data.txt"); err2 == nil {
		fmt.Printf("%s\n", data2)
	}
}
