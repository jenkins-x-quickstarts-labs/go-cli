package main

import (
	"os"

	"github.com/jenkins-x-labs/foo/cmd/root"
)

// Entrypoint for the command
func main() {
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
