package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen:err:", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn:err:", err)
			return
		}
		go Handel(conn)
	}
}

func Handel(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect success")

	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("connect Read err:", err)
			return
		}

		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		if "exit" == string(buf[:n-1]) {
			fmt.Println(addr, "exit")
			return
		}
		_, _ = conn.Write([]byte(string(buf[:n])))
	}
}
