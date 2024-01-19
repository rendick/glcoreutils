package main

import (
	"fmt"
	"os"
)

var (
	// Info
	rf_help = ` 
Usage: cwd [OPTION]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	rf_version = "0.1v"
)

func help() {
	fmt.Println("rf: missing file\nUsage: rf [OPTION or FILENAME]\nTry --help for more information")
}

func main() {
	if len(os.Args) != 3 {
		help()
		os.Exit(0)
	}

	filename := os.Args[1]
	file := os.Args[2]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(rf_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(rf_version)
	default:
		rename := os.Rename(filename, file)
		if rename == nil {
			fmt.Println("")
		} else {
			panic(rename)
		}
	}
}
