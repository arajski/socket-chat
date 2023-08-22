package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type Message struct {
	size   int
	client string
	data   string
}

func handleClient(conn *net.TCPConn, messages chan<- Message, clients map[string]*net.TCPConn) {
	clients[conn.RemoteAddr().String()] = conn
	defer delete(clients, conn.RemoteAddr().String())

	fmt.Printf("Connection with %s has been estabilished\n", conn.RemoteAddr())
	conn.Write([]byte("Connection estabilished! Welcome to the socket-chat!\n"))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil || n < 1 {
			break
		}

		msg := string(buf[:n])
		messages <- Message{size: n, client: conn.RemoteAddr().String(), data: msg}

		if msg == "exit" {
			conn.Close()
			break
		}
	}

	fmt.Printf("Connection with %s has been closed\n", conn.RemoteAddr())
}

func getMessages(messages <-chan Message, clients map[string]*net.TCPConn) {
	for {
		msg := <-messages
		fmt.Printf("Received %d bytes from %s\n", msg.size, msg.client)

		for _, conn := range clients {
			fmt.Fprintf(conn, "[%s]: %s", msg.client, msg.data)
		}
	}
}

func StartServer(host string, port int) {
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

	clients := make(map[string]*net.TCPConn)
	messages := make(chan Message)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(conn.(*net.TCPConn), messages, clients)
		go getMessages(messages, clients)
	}
}
