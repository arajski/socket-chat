package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var args = os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: showip [address]")
		return
	}

	var ips, err = net.LookupIP(os.Args[1])
	if err != nil {
		fmt.Printf("Error while parsing address")
		return
	}

	for _, ip := range ips {
		if len(ip) == net.IPv4len {
			fmt.Println("IPv4: ", ip)
		}
		if len(ip) == net.IPv6len {
			fmt.Println("IPv6: ", ip)
		}
	}

}
