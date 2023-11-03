package client

import (
	"github.com/arajski/socket-chat/pkg/client"
	"github.com/spf13/cobra"
)

var host string
var port int

var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Connect the socket-chat client",
	Long:  "Connects the socket-chat server to a given host and port",
	Run: func(cmd *cobra.Command, args []string) {
		client := client.Connect(host, port)
		client.Run()
	},
}

func init() {
	ClientCmd.Flags().StringVar(&host, "host", "", "A server hostname. Defaults to localhost")
	ClientCmd.Flags().IntVar(&port, "port", 3000, "A server port. Defaults to 3000")
}
