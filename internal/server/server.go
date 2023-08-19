package server

import (
	"fmt"
	"net"
	"os"
)

func Server(args []string) {
	if len(args) < 1 {
		fmt.Println("no port provider")
		fmt.Println("usage: sc server [port]")
		os.Exit(1)
	}

	port := &args[0]
	fmt.Printf("Starting up server on port %s...\n", *port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", *port)
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
			if err != nil {
				continue
			}

			if n < 1 {
				continue
			}

			fmt.Printf("Received %d bytes\n", n)
			msg := string(buf[:n])

			if msg == "exit" {
				conn.Close()
				break
			}
		}
		fmt.Printf("Connection with %s has been closed\n", conn.RemoteAddr())
	}
}
