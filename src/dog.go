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
Usage: dog [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	dog_version = "0.1v"

	// Color
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Red   = "\033[31m"
)

func help() {
	fmt.Println("dog: missing file\nUsage: dog: [OPTION of FILENAME]\nTry --help for more information")
}

func main() {
	if len(os.Args) != 2 {
		help()
		os.Exit(0)
	}

	filename := os.Args[1]
	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(dog_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(dog_version)
	default:
		file_count, err := os.Open(filename)
		if err != nil {
			os.Exit(0)
		}

		file, err := ioutil.ReadFile(filename)
		fmt.Println(string(file))
		scanner := bufio.NewScanner(file_count)
		scanner.Split(bufio.ScanLines)

		var count int
		for scanner.Scan() {
			count++
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf(Red+"Number of lines in file: "+Reset+"%d", count)
		}
	}
}
