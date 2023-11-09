package cmd

import (
	"github.com/arajski/socket-chat/pkg/client"
	"github.com/spf13/cobra"
)

var host string
var port int

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Connect the socket-chat client",
	Long:  "Connects the socket-chat server to a given host and port",
	Run:   startClient,
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringP("host", "H", "127.0.0.1", "A server hostname. Defaults to localhost")
	clientCmd.Flags().IntP("port", "P", 3000, "A server port. Defaults to 3000")
}

func startClient(cmd *cobra.Command, args []string) {

	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetInt("port")

	client := client.Connect(host, port)
	client.Run()
}
