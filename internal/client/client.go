package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type Client struct {
	conn    *net.TCPConn
	scanner *bufio.Scanner
}

func Connect(host string, port int) *Client {
	address := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))
	log.Printf("connecting client to %s...\n", address)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalln("invalid hostname:", address)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalln("could not connect to the host")
	}

	client := Client{conn, bufio.NewScanner(os.Stdin)}
	return &client
}

func (client *Client) Run() {
	go client.readMessages()

	for client.scanner.Scan() {
		client.conn.Write([]byte(client.scanner.Text()))

		if client.scanner.Text() == "exit" {
			os.Exit(0)
		}
	}
}

func (client *Client) readMessages() {
	for client.conn != nil {
		buf := make([]byte, 1024)
		response, err := client.conn.Read(buf)
		if err != nil {
			log.Fatalf("could not receive a response from server")
		}

		log.Println(string(buf[:response]))
	}
}
