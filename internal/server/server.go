package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func Server(host string, port int) {
	address := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))

	fmt.Printf("Starting up server on port %d...\n", port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Connection with %s has been estabilished\n", conn.RemoteAddr())
		conn.Write([]byte("Connection estabilished! Welcome to the socket-chat!\n"))

		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil || n < 1 {
				continue
			}

			fmt.Printf("Received %d bytes\n", n)
			msg := string(buf[:n])

			if msg == "exit" {
				conn.Close()
				break
			}

			fmt.Fprintf(conn, "[%s]: %s", conn.RemoteAddr(), msg)
		}

		fmt.Printf("Connection with %s has been closed\n", conn.RemoteAddr())
	}
}
