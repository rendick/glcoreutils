package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	user_help = `
Usage: user [OPTION]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
	`
	user_version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		user, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $1}' /etc/group").Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", user)
		os.Exit(0)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(user_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(user_version)
	}
}
