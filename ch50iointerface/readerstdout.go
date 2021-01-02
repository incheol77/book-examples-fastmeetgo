package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := "hello world\n"
	r := strings.NewReader(s)
	io.Copy(os.Stdout, r)
}
