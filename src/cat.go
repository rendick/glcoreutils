package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	cat_help = ` 
Usage: cat [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

--number or -n:	  counts the lines 

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	cat_version = "0.1v"

	Reset = "\033[0m"
	Bold  = "\033[1m"
	Red   = "\033[31m"
)

func help() {
	fmt.Println("cat: missing file\nUsage: cat [FILENAME or OPTION]\nTry --help for more information")
}

func cat_default() {
	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(cat_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(cat_version)
	default:
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			os.Exit(0)
		} else {
			fmt.Println(string(file))
		}
	}
}

func cat_count() {
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
		go cat_default()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	} else if len(os.Args) == 3 {
		go cat_count()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	} else {
		go help()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

}
