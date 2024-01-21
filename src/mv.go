package main

import (
	"fmt"
	"os"
	"time"
)

var (
	mv_help = ` 
Usage: mv [OPTION]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	mv_version = "0.1v"
)

func help() {
	fmt.Println("mv: missing file\nUsage: mv [OPTION or FILENAME]\nTry --help for more information")
}

func mv() {
	filename := os.Args[1]
	file := os.Args[2]

	mv := os.Rename(filename, file)
	if mv == nil {
		fmt.Println("")
	} else {
		panic(mv)
	}
}

func main() {
	if len(os.Args) != 3 {
		help()
		os.Exit(0)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(mv_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(mv_version)
	default:
		go mv()
		time.Sleep(time.Millisecond)
	}
}
