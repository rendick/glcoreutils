package main

import (
	"fmt"
	"os"
)

var (
	yes_help = ` 
Usage: yes [OPTION or TEXT]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	yes_version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		for {
			fmt.Println("yes")
		}
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(yes_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(yes_version)
	default:
		for {
			fmt.Println(filename)
		}
	}
}
