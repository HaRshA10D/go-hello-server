package cmd

import (
	"fmt"
	"github.com/HaRshA10D/go-hello-server/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Starts the http server",
	Long: "Starts the http server from port in config If not provided in flags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Http server will soon start....")
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
