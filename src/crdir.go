package main

import (
	"fmt"
	"log"
	"os"
)

var (
	// Info
	crdir_help = ` 
Usage: crdir [OPTION or DIR]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	crdir_version = "0.1v"
)

func help() {
	fmt.Println("crdir: missing file\nUsage: dog [OPTION or DIR]\nTry --help for more information")
}

func main() {
	if len(os.Args) != 2 {
		help()
		os.Exit(0)
	}

	filename := os.Args[1]
	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(crdir_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(crdir_version)
	default:
		err := os.Mkdir(filename, 0750)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		} else {
			fmt.Println(filename)
		}
	}
}
