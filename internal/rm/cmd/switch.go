package cmd

import (
	"fmt"
	"os"

	config "github.com/rendick/glcoreutils/internal/rm/settings"
)

func Switch() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command (e.g., --help or --version)")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "--help", "-h":
		// config.Help()
		Deletion()
	case "--version", "-v":
		config.Version()
	default:
		Deletion()
	}
}
