package main

import (
	"fmt"
	cmd "grpc_bri/pkg/cmd/server"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
