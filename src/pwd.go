package main

import (
	"fmt"
	"os"
	"time"
)

var (
	pwd_help = ` 
Usage: pwd [OPTION]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	pwd_version = "0.1v"
)

func pwd() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pwd)
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) != 2 {
		go pwd()
		time.Sleep(time.Millisecond)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(pwd_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(pwd_version)
	}
}
