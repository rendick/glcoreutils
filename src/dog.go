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
	"time"
)

var (
	// Info
	dog_help = ` 
Usage: dog [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

--number or -n:	  counts the lines 

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	dog_version = "0.2v"

	// Color
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Red   = "\033[31m"
)

func help() {
	fmt.Println("dog: missing file\nUsage: dog [FILENAME or OPTION]\nTry --help for more information")
}

func dog_default() {
	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(dog_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(dog_version)
	default:
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			os.Exit(0)
		} else {
			fmt.Println(string(file))
		}
	}
}

func dog_count() {
	filename := os.Args[1]
	args := os.Args[2]

	switch args {
	case "-n":
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
			fmt.Println(Red+"Number of lines in file: "+Reset, count)
		}
	}
}

func main() {
	if len(os.Args) == 2 {
		go dog_default()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	} else if len(os.Args) == 3 {
		go dog_count()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	} else {
		go help()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

}
