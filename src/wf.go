package main

import (
	"fmt"
	"os"
)

var (
	wf_help = `
Usage: wf [filename]

--help or -h:     help 
--version or -v:  version

GL coreutils online: <https://github.com/rendick/glcoreutils/>
	`
	version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: fc [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]

	if filename == "--help" || filename == "-h" {
		fmt.Println(wf_help)
	} else if filename == "--version" || filename == "-v" {
		fmt.Println(version)
	} else {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()

		fmt.Println("File written: ", filename)
	}

}
