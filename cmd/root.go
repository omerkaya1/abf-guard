package cmd

import (
	"github.com/omerkaya1/abf-guard/cmd/grpc"
	"github.com/spf13/cobra"
)

// RootCmd .
var RootCmd = &cobra.Command{
	Use:   "abf-guard",
	Short: "abf-guard is a small service that prevents brute force login attacks",
}

func init() {
	RootCmd.AddCommand(grpc.ClientRootCmd, grpc.ServerRootCmd)
}
