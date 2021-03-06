package main

import (
	"fmt"
	"net"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	go func(c net.Conn) {
		data := make([]byte, 4096)
		for cnt := 0; cnt < 10; cnt++ {
			n, err := c.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(data[:n]) + " I'm fine, and you?")
		}
	}(client)

	fmt.Scanln()
}
