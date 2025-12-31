// Package main provides the CLI entry point for Velocity projects.
// This file is the development CLI - use `go build ./cmd/velocity` to build it.
// Production builds should use `go build .` which builds only the server.
package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/velocitykode/velocity-cli/cli"

	// Import project packages to register with CLI
	_ "{{MODULE_NAME}}/bootstrap"
	_ "{{MODULE_NAME}}/database/migrations"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Run CLI
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
