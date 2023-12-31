package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of socket-chat",
	Long:  `This is a version of the socket-chat`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("socket-chat v0.1")
	},
}
