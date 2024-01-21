package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	id_help = `
Usage: id [OPTION or FILENAME]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
	`
	id_version = "0.1v"
)

func id() {
	id, err := exec.Command("sh", "-c", "awk -F: '$3==998 {print $0} $3==1000 {print $0}' /etc/group").Output()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s", id)
	}
}

func id_group() {
	id_group, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $3}' /etc/group").Output()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s", id_group)
	}
}

func id_groups() {
	id_groups, err := exec.Command("sh", "-c", "awk -F: '$3==998 {print $3} $3==1000 {print $3}' /etc/group").Output()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s", id_groups)
	}
}

func id_user() {
	id_user, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $3}' /etc/group").Output()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s", id_user)
	}
}

func main() {
	if len(os.Args) != 2 {
		go id()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(id_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(id_version)
	case "--group", "-group", "-g", "--g":
		go id_group()
		time.Sleep(time.Millisecond)
	case "--groups", "-groups", "-G", "--G":
		go id_groups()
		time.Sleep(time.Millisecond)
	case "--user", "-user", "-u", "--u":
		go id_user()
		time.Sleep(time.Millisecond)
	}
}
