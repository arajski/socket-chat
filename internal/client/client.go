package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func readMessages(conn *net.TCPConn) {
	for conn != nil {
		buf := make([]byte, 1024)
		response, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("could not receive a response from server")
		}

		log.Println(string(buf[:response]))
	}
}

func AttachClient(host string, port int) {
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

	go readMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		conn.Write([]byte(scanner.Text()))

		if scanner.Text() == "exit" {
			os.Exit(0)
		}
	}
}
