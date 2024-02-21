package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Switch() {
	version := flag.Bool("v", false, "version")
	flag.Parse()
	if *version {
		fmt.Println("sleep 1.0v")
		os.Exit(0)
	}
	Sleep()

}
