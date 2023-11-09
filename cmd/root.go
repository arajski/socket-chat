package cmd

import (
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
}

func Execute() error {
	return rootCmd.Execute()
}
