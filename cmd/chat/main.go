package main

import (
	"flag"
	"internal/client"
	"internal/server"
	"log"
	"os"
)

type Runner interface {
	Run()
}

var runner Runner

func init() {
	var host string
	var port int

	flagSet := flag.NewFlagSet("flags", flag.ExitOnError)
	flagSet.StringVar(&host, "hostname", "127.0.0.1", "Name of a host")
	flagSet.IntVar(&port, "port", 3000, "Running port")

	if len(os.Args) < 2 {
		log.Fatalln("expected 'server' or 'client' subcommands")
	}

	flagSet.Parse(os.Args[2:])

	switch os.Args[1] {
	case "server":
		runner = server.NewServer(host, port)
	case "client":
		runner = client.Connect(host, port)
	default:
		log.Fatalln("expected 'client' or 'server' subcommands")
		os.Exit(1)
	}
}

func main() {
	runner.Run()
}
