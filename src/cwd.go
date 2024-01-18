package main

import (
	"fmt"
	"os"
)

var (
	// Info
	cwd_help = ` 
Usage: cwd [OPTION]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	cwd_version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		} else {
			fmt.Println(cwd)
			os.Exit(0)
		}
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(cwd_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(cwd_version)
	}
}
