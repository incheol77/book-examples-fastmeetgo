package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile(
		"stringwriter.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	s := "hello world\n"
	r := strings.NewReader(s)

	w := bufio.NewWriter(file)

	w.ReadFrom(r)
	w.Flush()
}
