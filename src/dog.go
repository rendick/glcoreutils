package main

import (
	"fmt"
	"os"
)

func main() {
	x := os.Args[1]

	if x == "--help" {
		fmt.Println("ealjfkhl")
	}

	file, err := os.Open(x)
	if err != nil {
		fmt.Println("sklfjshlk")
	}
	defer file.Close()

}
