package main

import (
	"fmt"
	"os"
	"time"
)

var (
	mkdir_help = ` 
Usage: mkdir [OPTION or DIR]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	mkdir_version = "0.1v"
)

func help() {
	fmt.Println("mkdir: missing file\nUsage: mkdir [OPTION or DIR]\nTry --help for more information")
}

func mkdir() {
	filename := os.Args[1]

	err := os.Mkdir(filename, 0750)
	if err != nil && !os.IsExist(err) {
		panic(err)
	} else {
		fmt.Println(filename)
	}
}

func main() {
	if len(os.Args) != 2 {
		go help()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	} else {
		go mkdir()
		time.Sleep(time.Millisecond)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(mkdir_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(mkdir_version)
	default:
		go mkdir()
		time.Sleep(time.Millisecond)
	}
}
