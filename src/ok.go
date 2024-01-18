package main

import (
	"fmt"
	"os"
)

var (
	ok_help = ` 
Usage: ok [OPTION or TEXT]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	ok_version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		for {
			fmt.Println("ok")
		}
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(ok_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(ok_version)
	default:
		for {
			fmt.Println(filename)
		}
	}
}
