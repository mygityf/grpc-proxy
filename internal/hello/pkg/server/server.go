package main

import (
	"fmt"
	"grpc-proxy/internal/hello/pkg/cmd"
	"os"
)

var (
	// root cmd
	rootCmd = cmd.NewHelloCommand()
)

// main
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("rootCmd err:%t\n", err)
		os.Exit(1)
	}
}
