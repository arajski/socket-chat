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

type ChatServer struct {
	listener *net.TCPListener
	messages chan Message
	signals  chan os.Signal
	clients  map[string]*net.TCPConn
}

func NewServer(host string, port int) *ChatServer {
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

	server := ChatServer{listener, make(chan Message, 100), make(chan os.Signal, 2), make(map[string]*net.TCPConn)}
	return &server
}

func (server ChatServer) Run() {
	signal.Notify(server.signals, os.Interrupt, syscall.SIGTERM)
	go server.handleMessages()
	go server.handleSignals()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go server.handleClient(conn.(*net.TCPConn))
	}
}

func (server ChatServer) handleClient(conn *net.TCPConn) {
	server.clients[conn.RemoteAddr().String()] = conn
	defer conn.Close()
	defer delete(server.clients, conn.RemoteAddr().String())

	log.Printf("connection with %s has been estabilished\n", conn.RemoteAddr())
	log.Printf("number of clients: %d\n", len(server.clients))

	conn.Write([]byte("connection estabilished! Welcome to the socket-chat!\n"))

	buf := make([]byte, 128)
	for {
		n, err := conn.Read(buf)
		if err != nil || n < 1 {
			break
		}

		msg := string(buf[:n])
		server.messages <- Message{size: n, client: conn.RemoteAddr().String(), data: msg}

		log.Printf("received %d bytes from %s\n", n, conn.RemoteAddr().String())

		if msg == "exit" {
			break
		}
	}

	log.Printf("sonnection with %s has been closed\n", conn.RemoteAddr())
}

func (server ChatServer) handleMessages() {
	for {
		msg := <-server.messages

		for _, conn := range server.clients {
			go func(conn *net.TCPConn) {
				log.Printf("sending %d bytes to %s\n", msg.size, conn.RemoteAddr().String())
				fmt.Fprintf(conn, "[%s]:%s", msg.client, msg.data)
			}(conn)
		}
	}
}

func (server ChatServer) handleSignals() {
	<-server.signals
	log.Fatal("shutting down...")
}
