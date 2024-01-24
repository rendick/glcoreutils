package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	sleep_help = ` 
Usage: sleep [OPTION or NUMBER]

--second or -s:   for seconds
--minute or -m:   for minutes
--hour or -h:     for hours
--day or -d:	  for days

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	sleep_version = "0.1v"
)

func help() {
	fmt.Println("sleep: missing operand\nUsage: sleep [NUMBER or OPTION]\nTry --help for more information")
}

func sleep_minute() {
	args := os.Args[2]

	switch args {
	case "-m":
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		} else {
			time.Sleep(time.Duration(i) * time.Minute)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		go help()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	} else if len(os.Args) == 2 {

		switch os.Args[1] {
		case "--help", "-help", "-h", "--h":
			fmt.Println(sleep_help)
		case "--version", "-version", "-v", "--v":
			fmt.Println(sleep_version)
		default:
			i, err := strconv.Atoi(os.Args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			} else {
				time.Sleep(time.Duration(i) * time.Second)
			}
		}
	} else if len(os.Args) == 3 {
		sleep_minute()
	}
}
