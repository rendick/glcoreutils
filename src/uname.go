// TODO: provide Android support

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

-a or --all:                print all information, in the following order:

-s or --kernel-name:        print the kernel name
-n or --nodename:           print the network node hostname
-r or --kernel-release:     print the kernel release
-v or --kernel-version:     print the kernel version
-m or --machine:            print the machine hardware name
-p or --processor:          print the processor type (non-portable)
-i or --hardware-platform:  print the hardware platform (non-portable)
-o or --operating-system:   print the operating system

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
	} else {
		fmt.Println(nodename)
	}
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
	kernel_version, err := os.ReadFile("/proc/sys/kernel/version")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(strings.TrimSpace(string(kernel_version)))
	}
}

func uname_machine() {
	uname_machine, err := os.ReadFile("/proc/sys/kernel/arch")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(strings.TrimSpace(string(uname_machine)))
	}
}

func uname_processor() {
	cpu, err := os.Open("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}

	defer cpu.Close()

	scanner := bufio.NewScanner(cpu)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "model name") {
			fields := strings.Split(line, ":")
			if len(fields) == 2 {
				modelName := strings.TrimSpace(fields[1])
				fmt.Println(modelName)
				return
			}
		}
	}
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
			fmt.Println(fields[0])
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
	} else if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--all", "-all", "-a", "--a":
			uname()
			uname_nodename()
			uname_kernel_release()
			uname_kernel_version()
			uname_machine()
			uname_processor()
			uname_os()

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
			go uname_kernel_version()
			time.Sleep(time.Millisecond)
		case "--machine", "-machine", "-m", "--m":
			go uname_machine()
			time.Sleep(time.Millisecond)

		case "--processor", "-processor", "-p", "--p":
			go uname_processor()
			time.Sleep(time.Millisecond)

		case "--hardware-platform", "-hardware-platform", "-i", "--i":
			fmt.Println("Soon.")

		case "--operating-system", "-operating-system", "-o", "--o":
			go uname_os()
			time.Sleep(time.Millisecond)

		default:
			fmt.Printf("uname: invalid option %s \nTry --help for more information.", os.Args[1])
		}
	} else {
		fmt.Println("Error!")
	}
}
