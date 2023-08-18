package main

import (
	"fmt"
	"net"
	"os"
)

func handleServer(args []string) {
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

		conn.Write([]byte("Connection estabilished! Welcome to the socket-chat!\n"))
		conn.Write([]byte("To list channels, type 'list'\n"))
		conn.Write([]byte("To join a channel, type 'join'\n"))
		conn.Write([]byte("To create a channel, type 'create'\n"))

		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("could not receive a response from client")
			}
			msg := string(buf[:n])
			fmt.Println(msg)
			if msg == "exit" {
				conn.Close()
				break
			}
		}
	}
}

func handleClient(args []string) {
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

	conn.Write([]byte("test"))
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no parameters provided")
		fmt.Println("usage: sc client [hostname]")
		os.Exit(1)
	}

	switch args[0] {
	case "client":
		handleClient(args[1:])
	case "server":
		handleServer(args[1:])
	default:
		fmt.Println("unknown command")
		fmt.Println("usage: sc [server/client]")
		os.Exit(1)
	}
}
