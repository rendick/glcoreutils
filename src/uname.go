// TODO: processor, kernel-version,

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	uname_help = ` 
Usage: uname [OPTION]

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	uname_version = "0.1v"
)

func uname() {
	uname, err := os.Open("/proc/version")
	if err != nil {
		panic(err)
	}
	defer uname.Close()
	r := bufio.NewScanner(uname)
	for r.Scan() {
		fields := strings.Fields(r.Text())
		if len(fields) > 0 {
			fmt.Println(fields[0])
		}

		if err != nil {
			fmt.Println("Error reading file: ", err)
		}
	}
}

func uname_nodename() {
	nodename, err := os.Hostname()
	if err != nil {
		fmt.Println(nodename, err)
	}
	fmt.Println(nodename)
}

func uname_kernel_release() {
	uname_kernel_release, err := os.Open("/proc/version")
	if err != nil {
		panic(err)
	}
	defer uname_kernel_release.Close()
	r := bufio.NewScanner(uname_kernel_release)
	for r.Scan() {
		fields := strings.Fields(r.Text())
		if len(fields) > 0 {
			fmt.Println(fields[2])
		}

		if err != nil {
			fmt.Println("Error reading file: ", err)
		}
	}
}

func uname_kernel_version() {
	fmt.Println("Soon.")
}

func uname_machine() {
	uname_machine, err := os.ReadFile("/proc/sys/kernel/arch")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.TrimSpace(string(uname_machine)))
}

func uname_processor() {
	uname_processor := "Soon."
	fmt.Printf("%s", uname_processor)
}

func uname_os() {
	uname_os, err := os.Open("/proc/version")
	if err != nil {
		os.Exit(0)
	}
	defer uname_os.Close()
	r := bufio.NewScanner(uname_os)
	for r.Scan() {
		fields := strings.Fields(r.Text())
		if len(fields) > 0 {
			if fields[0] == "Linux" {
				fmt.Printf("GNU/%s", fields[0])
			} else if fields[0] == "Android" {
				fmt.Println("Android")
			}
		}
		if err != nil {
			fmt.Println("Error reading file: ", err)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		go uname()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

	filename := os.Args[1]

	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(uname_help)
	case "--version", "-version":
		fmt.Println(uname_version)
	case "--kernel-name", "-kernel-name", "-s", "--s":
		go uname()
		time.Sleep(time.Millisecond)
	case "--nodename", "-nodename", "-n", "--n":
		go uname_nodename()
		time.Sleep(time.Millisecond)
	case "--kernel-release", "-kernel-release", "-r", "--r":
		go uname_kernel_release()
		time.Sleep(time.Millisecond)
	case "--kernel-version", "-kernel-version", "-v", "--v":
		fmt.Println("Soon.")
	case "--machine", "-machine", "-m", "--m":
		go uname_machine()
		time.Sleep(time.Millisecond)
	case "--processor", "-processor", "-p", "--p":
		go uname_processor()
		time.Sleep(time.Millisecond)
	case "--operating-system", "-operating-system", "-o", "--o":
		go uname_os()
		time.Sleep(time.Millisecond)
	}
}
