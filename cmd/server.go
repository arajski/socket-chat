package cmd

import (
	"github.com/arajski/socket-chat/pkg/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the socket-chat server",
	Long:  "Starts the socket-chat server on a given port",
	Run:   startServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringP("host", "H", "127.0.0.1", "A server hostname. Defaults to localhost")
	serverCmd.Flags().IntP("port", "P", 3000, "A server port. Defaults to 3000")
}

func startServer(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetInt("port")

	server := server.NewServer(host, port)
	server.Run()
}
