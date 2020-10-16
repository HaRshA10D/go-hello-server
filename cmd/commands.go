package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "go-hello-server",
	Short: "A simple http server application",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("Error occured in executing root command, \n Description: %s", err)
		panic(err)
	}
}
