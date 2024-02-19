package main

import (
	"fmt"
	"os"
	"strings"
)

func yes() {
	for {
		if len(os.Args) == 1 {
			fmt.Println("y")
		} else {
			fmt.Println(strings.Join(os.Args[1:], " "))
		}
	}
}

func main() {
	yes()
}
