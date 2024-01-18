package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	usrid_help = `
Usage: wf [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
	`
	usrid_version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		usrid, err := exec.Command("sh", "-c", "awk -F: '$3==998 {print $0} $3==1000 {print $0}' /etc/group").Output()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", usrid)
		}
		os.Exit(0)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(usrid_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(usrid_version)
	case "--group", "-group", "-g", "--g":
		group_usrid, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $3}' /etc/group").Output()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", group_usrid)
		}
	case "--groups", "-groups", "-G", "--G":
		groups_usrid, err := exec.Command("sh", "-c", "awk -F: '$3==998 {print $3} $3==1000 {print $3}' /etc/group").Output()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", groups_usrid)
		}
	case "--user", "-user", "-u", "--u":
		user_usrid, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $3}' /etc/group").Output()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", user_usrid)
		}
	}
}
