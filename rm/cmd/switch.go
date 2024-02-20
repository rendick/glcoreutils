// switch.go

package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func Switch() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command (e.g., --help or --version)")
		os.Exit(0)
	}

	dir := flag.String("r", "", "folder")
	version := flag.Bool("v", false, "Print version")
	clean := flag.Bool("c", false, "clean")
	flag.Parse()

	if *version {
		// Version()
		fmt.Println("vers")
		os.Exit(0)
	}

	if *clean {
		Clean() // Call Clean function without passing any arguments
		os.Exit(0)
	}

	if *dir == "" {
		log.Fatal("Please provide a folder to delete using the -r flag.")
	}

	switch os.Args[1] {
	default:
		Deletion(*dir)
	}
}
