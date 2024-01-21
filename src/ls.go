package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	ls_help = ` 
Usage: ls [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	ls_version = "0.1v"
)

func ls() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ls, err := os.ReadDir(strings.TrimSpace(string(dir)))
	if err != nil {
		panic(err)
	}

	for _, files := range ls {
		fmt.Println(files.Name())
	}
}

func main() {
	if len(os.Args) != 2 {
		go ls()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

	filename := os.Args[1]
	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(ls_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(ls_version)
	}
}
