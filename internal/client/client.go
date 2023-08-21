package client

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func Client(host string, port int) {
	address := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))
	fmt.Printf("Connecting client to %s...\n", address)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("invalid hostname:", address)
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
