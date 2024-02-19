package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	Reset = "\033[0m"
)

func yes() {

	color_arr := map[int]string{
		1: "\033[31m",
		2: "\033[32m",
		3: "\033[33m",
		4: "\033[35m",
	}

	gen := rand.Perm(4)

	for {
		for _, index := range gen {
			color := color_arr[index+1]

			if len(os.Args) == 1 {
				fmt.Println(color + "y" + Reset)
			} else {
				fmt.Println(color + strings.Join(os.Args[1:], " ") + Reset)
			}
		}
	}
}

func main() {
	yes()
}
