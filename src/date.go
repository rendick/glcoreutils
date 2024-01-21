package main

import (
	"fmt"
	"os"
	"time"
)

var (
	date_help = `
Usage: date [OPTION]

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
	date_version = "0.1v"
)

func date() {
	fmt.Println(time.Now().Format(time.UnixDate))
}

func date_day() {
	fmt.Println(time.Now().Weekday())
}

func date_month() {
	fmt.Println(time.Now().Month())
}

func date_year() {
	fmt.Println(time.Now().Year())
}

func date_hour() {
	fmt.Println(time.Now().Hour())
}

func date_century() {
	fmt.Println((time.Now().Year() + 99) / 100)
}

func date_week() {
	fmt.Println(time.Now().ISOWeek())
}

func main() {
	if len(os.Args) != 2 {
		go date()
		time.Sleep(time.Millisecond)
		os.Exit(0)
	}

	filename := os.Args[1]
	switch filename {
	case "--help", "-help", "-h", "--h":
		fmt.Println(date_help)

	case "--version", "-version", "-v", "--v":
		fmt.Println(date_version)

	case "--day", "-day", "-d", "--d":
		go date_day()
		time.Sleep(time.Millisecond)

	case "--month", "-month", "-m", "--m":
		go date_month()
		time.Sleep(time.Millisecond)

	case "--year", "-year", "-y", "--y":
		go date_year()
		time.Sleep(time.Millisecond)

	case "--hour", "-hour", "-H", "--H":
		go date_hour()
		time.Sleep(time.Millisecond)

	case "--century", "-century", "-c", "--c":
		go date_century()
		time.Sleep(time.Millisecond)

	case "--week", "-week", "-w", "--w":
		go date_week()
		time.Sleep(time.Millisecond)
	}
}
