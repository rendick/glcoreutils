package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	dog_help = ` 
Usage: dog [filename]

--help or -h:     help 
--version or -v:  version

GL coreutils online: <https://github.com/rendick/glcoreutils/>
`
	version = "0.1v"
)

func help() {
	fmt.Println("dog: missing file\nUsage: dog: [filename]\nTry --help for more information")
}

func main() {
	if len(os.Args) != 2 {
		help()
		os.Exit(0)
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(dog_help)
	} else if os.Args[1] == "--version" || os.Args[1] == "-v" {
		fmt.Println(version)
	}

	filename := os.Args[1]

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(0)
	}
	fmt.Println(string(file))
}
