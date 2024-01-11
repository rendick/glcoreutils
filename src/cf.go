package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./fc [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("%s", filename)
}
