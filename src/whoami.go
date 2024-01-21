package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"
)

var (
	whoami_help = `
Usage: whoami [OPTION]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
	`
	whoami_version = "0.1v"
)

func whoami() {
	whoami, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(whoami)
}

func main() {
	if len(os.Args) != 2 {
		// user, err := exec.Command("sh", "-c", "awk -F: '$3==1000 {print $1}' /etc/group").Output()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Printf("%s", user)

		go whoami()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(whoami_help)
	case "--version", "-version", "-v", "--v":
		fmt.Println(whoami_version)
	}
}
