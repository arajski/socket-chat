package cmd

import (
	"github.com/arajski/socket-chat/cmd/client"
	"github.com/arajski/socket-chat/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "socket-chat",
	Short: "A terminal chat application written in Go",
	Long: `socket-chat is a terminal chat application
	which utilizes web sockets`,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(client.ClientCmd)
	rootCmd.AddCommand(server.ServerCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
