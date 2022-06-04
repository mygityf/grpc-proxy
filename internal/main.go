package main

import (
	"fmt"
	"grpc-proxy/internal/proxy/pkg/cmd"
	"os"
)

var (
	// root cmd
	rootCmd = cmd.NewGrpcProxyCommand()
)

// main
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("rootCmd err:%t\n", err)
		os.Exit(1)
	}
}
