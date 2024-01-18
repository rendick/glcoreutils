// TODO: create a colorful output to display the type

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	// Info
	vd_help = ` 
Usage: vd [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	vd_version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		vd, err := os.ReadDir(strings.TrimSpace(string(dir)))
		if err != nil {
			panic(err)
		}

		for _, files := range vd {
			fmt.Println(files.Name())
		}
		os.Exit(0)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(vd_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(vd_version)
	}
}
