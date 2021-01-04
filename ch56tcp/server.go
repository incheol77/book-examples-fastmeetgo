package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("server : " + string(data[:n]))
		c.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	lsn, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lsn.Close()

	for {
		conn, err := lsn.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}
