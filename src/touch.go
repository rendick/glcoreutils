package main

import (
	"fmt"
	"os"
	"time"
)

var (
	touch_help = `
Usage: touch [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
	`
	touch_version = "0.1v"
)

func help() {
	fmt.Println("touch: missing file\nUsage: touch [OPTION of FILENAME]\nTry --help for more information")
}

func touch() {
	filename := os.Args[1]
	touch, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(0)
	} else {
		defer touch.Close()
		fmt.Println("File written: ", filename)
	}
}

func main() {
	filename := os.Args[1]

	if len(os.Args) != 2 {
		go help()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(touch_help)

	case "--version", "-version", "-v", "--v":
		fmt.Println(touch_version)

	default:
		go touch()
		time.Sleep(time.Millisecond)
	}
}
