package main

import (
	"fmt"
	"internal/client"
	"internal/server"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no parameters provided")
		fmt.Println("usage: sc client [hostname]")
		os.Exit(1)
	}

	switch args[0] {
	case "client":
		client.Client(args[1:])
	case "server":
		server.Server(args[1:])
	default:
		fmt.Println("unknown command")
		fmt.Println("usage: sc [server/client]")
		os.Exit(1)
	}
}
