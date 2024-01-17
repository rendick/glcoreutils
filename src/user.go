package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		user, err := exec.Command("sh", "-c", "awk -F: '$1==\"rendick\" {print $1}' /etc/group").Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", user)
		os.Exit(0)
	}

	filename := os.Args[1]
	if filename == "--help" || filename == "-h" {
		fmt.Println(user_help)
	} else if filename == "--version" || filename == "-v" {
		fmt.Println(user_version)
	}

}
