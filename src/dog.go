// {{ TODO:
// 1. `-n` flag for counting lines
// 2.
// }

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	// Info
	dog_help = ` 
Usage: dog [filename]

--help or -h:     help 
--version or -v:  version

GL coreutils online: <https://github.com/rendick/glcoreutils/>
`

	// Color
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Red   = "\033[31m"
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
		fmt.Println("0.1v")
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
		fmt.Printf(Red+"Number of lines in file: "+Reset+"%d", count)
	}
}
