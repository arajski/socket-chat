package client

import (
	"fmt"
	"net"
	"os"
)

func Client(args []string) {
	if len(args) < 1 {
		fmt.Println("hostname not provided")
		fmt.Println("usage: sc client [hostname:port]")
		os.Exit(1)
	}

	hostname := &args[0]
	fmt.Printf("Connecting client to %s...\n", *hostname)

	tcpAddr, err := net.ResolveTCPAddr("tcp", *hostname)
	if err != nil {
		fmt.Println("invalid hostname:", *hostname)
		os.Exit(0)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("could not connect to the host")
		os.Exit(0)
	}

	buf := make([]byte, 1024)
	response, err := conn.Read(buf)
	if err != nil {
		fmt.Println("could not receive a response from server")
	}

	fmt.Println(string(buf[:response]))

	for {
		var msg string
		fmt.Scanln(&msg)

		conn.Write([]byte(msg))

		if msg == "exit" {
			os.Exit(0)
		}
	}
}
