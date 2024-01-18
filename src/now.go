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
	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(now_help)

	case "--version", "-version", "-v", "--v":
		fmt.Println(now_version)

	case "--day", "-day", "-d", "--d":
		fmt.Println(time.Now().Weekday())

	case "--month", "-month", "-m", "--m":
		fmt.Println(time.Now().Month())

	case "--year", "-year", "-y", "--y":
		fmt.Println(time.Now().Year())

	case "--hour", "-hour", "-H", "--H":
		fmt.Println(time.Now().Hour())

	case "--century", "-century", "-c", "--c":
		fmt.Println((time.Now().Year() + 99) / 100)

	case "--week", "-week", "-w", "--w":
		fmt.Println(time.Now().ISOWeek())
	}
}
