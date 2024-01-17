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

	if filename == "--help" || filename == "-help" || filename == "-h" || filename == "--h" {
		fmt.Println(usrid_help)
	} else if filename == "--version" || filename == "-version" || filename == "-v" || filename == "--v" {
		fmt.Println(usrid_version)
	} else if filename == "--group" || filename == "-group" || filename == "-g" || filename == "--g" {
		group_usrid, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $3}' /etc/group").Output()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", group_usrid)
		}
	} else if filename == "--groups" || filename == "-groups" || filename == "-G" || filename == "--G" {
		groups_usrid, err := exec.Command("sh", "-c", "awk -F: '$3==998 {print $3} $3==1000 {print $3}' /etc/group").Output()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", groups_usrid)
		}
	} else if filename == "--user" || filename == "-user" || filename == "-u" || filename == "--u" {
		user_usrid, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $3}' /etc/group").Output()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", user_usrid)
		}
	}
}
