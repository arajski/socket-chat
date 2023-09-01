package server

import (
	"fmt"
	"log"
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

	log.Printf("connection with %s has been estabilished\n", conn.RemoteAddr())
	log.Printf("number of clients: %d\n", len(clients))

	conn.Write([]byte("connection estabilished! Welcome to the socket-chat!\n"))

	buf := make([]byte, 128)
	for {
		n, err := conn.Read(buf)
		if err != nil || n < 1 {
			break
		}

		msg := string(buf[:n])
		in <- Message{size: n, client: conn.RemoteAddr().String(), data: msg}

		log.Printf("received %d bytes from %s\n", n, conn.RemoteAddr().String())

		if msg == "exit" {
			conn.Close()
			break
		}
	}

	log.Printf("sonnection with %s has been closed\n", conn.RemoteAddr())
}

func handleMessages(in <-chan Message, clients map[string]*net.TCPConn) {
	for {
		msg := <-in

		for _, conn := range clients {
			go func(conn *net.TCPConn) {
				log.Printf("sending %d bytes to %s\n", msg.size, conn.RemoteAddr().String())
				fmt.Fprintf(conn, "[%s]:%s", msg.client, msg.data)
			}(conn)
		}
	}
}

func handleSignals(signals <-chan os.Signal) {
	<-signals
	log.Fatal("shutting down...")
}

func StartServer(host string, port int) {
	messages := make(chan Message, 100)
	signals := make(chan os.Signal, 2)
	clients := make(map[string]*net.TCPConn)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	address := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))

	log.Printf("started up server on port %d\n", port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	go handleMessages(messages, clients)
	go handleSignals(signals)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(clients, conn.(*net.TCPConn), messages)
	}
}
