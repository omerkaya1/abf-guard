package cmd

import "github.com/spf13/cobra"

// RootCmd is the main command to start the application
var RootCmd = &cobra.Command{
	Use:   "abf-guard",
	Short: "abf-guard is a small service that prevents brute force login attacks",
}

func init() {
	RootCmd.AddCommand(ClientRootCmd, ServerRootCmd)
}
