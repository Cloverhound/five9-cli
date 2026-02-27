package main

import (
	"os"

	"github.com/Cloverhound/five9-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
