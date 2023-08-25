package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Message struct {
	size   int
	client string
	data   string
}

func handleClient(clients map[string]*net.TCPConn, conn *net.TCPConn, in chan<- Message) {
	clients[conn.RemoteAddr().String()] = conn
	defer delete(clients, conn.RemoteAddr().String())

	fmt.Printf("[INFO] Connection with %s has been estabilished\n", conn.RemoteAddr())
	fmt.Printf("[INFO] Number of clients: %d\n", len(clients))

	conn.Write([]byte("[INFO] Connection estabilished! Welcome to the socket-chat!\n"))

	buf := make([]byte, 128)
	for {
		n, err := conn.Read(buf)
		if err != nil || n < 1 {
			break
		}

		msg := string(buf[:n])
		in <- Message{size: n, client: conn.RemoteAddr().String(), data: msg}

		fmt.Printf("[INFO] Received %d bytes from %s\n", n, conn.RemoteAddr().String())

		if msg == "exit" {
			conn.Close()
			break
		}
	}

	fmt.Printf("[INFO] Connection with %s has been closed\n", conn.RemoteAddr())
}

func handleMessages(in <-chan Message, clients map[string]*net.TCPConn) {
	for {
		msg := <-in

		for _, conn := range clients {
			go func(conn *net.TCPConn) {
				fmt.Printf("[INFO] Sending %d bytes to %s\n", msg.size, conn.RemoteAddr().String())
				fmt.Fprintf(conn, "[%s]:%s", msg.client, msg.data)
			}(conn)
		}
	}
}

func handleSignals(signals <-chan os.Signal) {
	<-signals
	fmt.Println("[ERROR] Shutting down...")
	os.Exit(1)
}

func StartServer(host string, port int) {
	messages := make(chan Message, 100)
	signals := make(chan os.Signal, 2)
	clients := make(map[string]*net.TCPConn)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	address := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))

	fmt.Printf("[INFO] Starting up server on port %d...\n", port)

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

	go handleMessages(messages, clients)
	go handleSignals(signals)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(clients, conn.(*net.TCPConn), messages)
	}
}
