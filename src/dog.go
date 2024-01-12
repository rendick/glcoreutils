package main

import (
	"bufio"
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
		// Count lines
		file_count, err := os.Open(filename)
		if err != nil {
			os.Exit(0)
		}

		// Read file
		file, err := ioutil.ReadFile(filename)
		fmt.Println(string(file))
		scanner := bufio.NewScanner(file_count)
		scanner.Split(bufio.ScanLines)

		// getting the number of lines
		var count int
		for scanner.Scan() {
			count++
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Number of lines in file:", count)
	}
}
