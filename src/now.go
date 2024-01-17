package main

import (
	"fmt"
	"os"
	"time"
)

var (
	// Info
	now_help = ` 
Usage: now [OPTION]

-d or --day: prints the full name of the weekday (e.g., Sunday)
-m or --month: prints the full name of the month (e.g,, January)
-y or --year: prints the year (e.g., 1984)
-H or -hour: prints an hour (e.g., 24)
-c or --century: prints the current century (e.g., 21)
-w or --week: prints the week number (e.g., 2017 18)

--help or -h:     help 
--version or -v:  version

GL coreutils: <https://github.com/rendick/glcoreutils/>
`
	now_version = "0.1v"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(time.Now().Format(time.UnixDate))
		os.Exit(0)
	}

	filename := os.Args[1]

	if filename == "--help" || filename == "-help" || filename == "-h" || filename == "--h" {
		fmt.Println(now_help)
	} else if filename == "--version" || filename == "-version" || filename == "-v" || filename == "--v" {
		fmt.Println(now_version)
	} else if filename == "--day" || filename == "-day" || filename == "-d" || filename == "-d" {
		fmt.Println(time.Now().Weekday())
	} else if filename == "--month" || filename == "-month" || filename == "-m" || filename == "--m" {
		fmt.Println(time.Now().Month())
	} else if filename == "--year" || filename == "-year" || filename == "-y" || filename == "--y" {
		fmt.Println(time.Now().Year())
	} else if filename == "--hour" || filename == "-hour" || filename == "-H" || filename == "--H" {
		fmt.Println(time.Now().Hour())
	} else if filename == "--century" || filename == "-century" || filename == "-c" || filename == "--c" {
		fmt.Println((time.Now().Year() + 99) / 100)
	} else if filename == "--week" || filename == "-week" || filename == "-w" || filename == "--w" {
		fmt.Println(time.Now().ISOWeek())
	}

}
