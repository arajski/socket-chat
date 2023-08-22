package main

import (
	"flag"
	"fmt"
	"internal/client"
	"internal/server"
	"os"
)

type Mode int32

const (
	Server Mode = iota
	Client
)

var (
	port int
	host string
	mode Mode
)

func init() {
	serverFlagSet := flag.NewFlagSet("server", flag.ExitOnError)
	serverFlagSet.StringVar(&host, "hostname", "127.0.0.1", "Name of a server's host")
	serverFlagSet.IntVar(&port, "port", 3000, "Server's running port")

	clientFlagSet := flag.NewFlagSet("client", flag.ExitOnError)
	clientFlagSet.StringVar(&host, "hostname", "127.0.0.1", "Name of a server's host")
	clientFlagSet.IntVar(&port, "port", 3000, "Server's running port")

	if len(os.Args) < 2 {
		fmt.Println("expected 'server' or 'client' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		serverFlagSet.Parse(os.Args[2:])
		mode = Server
	case "client":
		clientFlagSet.Parse(os.Args[2:])
		mode = Client
	default:
		fmt.Println("expected 'client' or 'server' subcommands")
		os.Exit(1)
	}
}

func main() {
	switch mode {
	case Server:
		server.HandleServer(host, port)
	case Client:
		client.HandleClient(host, port)
	default:
		fmt.Println("socket chat mode is not defined")
		os.Exit(1)
	}
}
