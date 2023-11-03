package server

import (
	"github.com/arajski/socket-chat/pkg/server"
	"github.com/spf13/cobra"
)

var host string
var port int

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the socket-chat server",
	Long:  "Starts the socket-chat server on a given port",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer(host, port)
		server.Run()
	},
}

func init() {
	ServerCmd.Flags().StringVar(&host, "host", "", "A server hostname. Defaults to localhost")
	ServerCmd.Flags().IntVar(&port, "port", 3000, "A server port. Defaults to 3000")
}
