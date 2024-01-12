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

	filename := os.Args[1]

	if filename == "--help" || filename == "-h" {
		fmt.Println(dog_help)
	} else if filename == "--version" || filename == "-v" {
		fmt.Println(version)
	} else {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(string(file))

	}
}
