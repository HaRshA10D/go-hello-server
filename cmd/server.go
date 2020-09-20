package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Starts the http server",
	Long: "Starts the http server from port in config If not provided in flags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Http server will soon start....")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
