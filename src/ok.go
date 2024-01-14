package main

import (
	"fmt"
	"os"
)

var (
	ok_help = ` 
Usage: ok [text]

--help or -h:     help 
--version or -v:  version

GL coreutils online: <https://github.com/rendick/glcoreutils/>
`
)

func main() {
	if len(os.Args) != 2 {
		for {
			fmt.Println("ok")

		}
	}
	x := os.Args[1]

	if x == "--help" || x == "-h" {
		fmt.Println(ok_help)
	} else if x == "--version" || x == "-v" {
		fmt.Println("0.1v")
	} else {
		for {
			fmt.Println(x)
		}
	}
}
