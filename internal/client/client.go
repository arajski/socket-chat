package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func readMessages(conn *net.TCPConn) {
	for conn != nil {
		buf := make([]byte, 1024)
		response, err := conn.Read(buf)
		if err != nil {
			fmt.Println("could not receive a response from server")
		}

		fmt.Println(string(buf[:response]))
	}
}

func AttachClient(host string, port int) {
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

	go readMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		conn.Write([]byte(scanner.Text()))

		if scanner.Text() == "exit" {
			os.Exit(0)
		}
	}
}
