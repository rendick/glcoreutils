package main

import (
	"fmt"
	"os"
)

var (
	wf_help = `
Usage: wf [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
	`
	wf_version = "0.1v"
)

func help() {
	fmt.Println("wf: missing file\nUsage: wf [OPTION of FILENAME]\nTry --help for more information")
}

func main() {
	if len(os.Args) != 2 {
		help()
		os.Exit(0)
	}

	filename := os.Args[1]

	if filename == "--help" || filename == "-h" {
		fmt.Println(wf_help)
	} else if filename == "--version" || filename == "-v" {
		fmt.Println(wf_version)
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
