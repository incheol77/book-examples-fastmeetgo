package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
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
			fmt.Println("Read goroutine : " + string(data[:n]))
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for cnt := 0; cnt < 10; cnt++ {
			s := "Write goroutine : hello " + strconv.Itoa(i)
			c.Write([]byte(s))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Complete say hello")
			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	fmt.Scanln()
}
