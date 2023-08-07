package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: client [hostname]")
	}

	addr := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println("invalid hostname:", os.Args[1])
		os.Exit(0)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("could not connect to the host")
		os.Exit(0)
	}

	conn.Write([]byte("test"))

	response, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("could not receive a response from server")
	}

	fmt.Println(string(response))
}
