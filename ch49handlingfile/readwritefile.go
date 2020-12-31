package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n := 0
	s := "안녕하세요"
	n, err = file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, " bytes save completed")

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())
	file.Seek(0, os.SEEK_SET)
	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, " bytes read completed")
	fmt.Println(string(data))
}
