package main

import (
	"fmt"
	"math/rand"
)

var (
	Reset = "\033[0m"
)

func main() {
	arr1 := map[int]string{
		1: "\033[31m",
		2: "\033[32m",
	}

	gen := rand.Perm(2)

	for _, index := range gen {
		color, exists := arr1[index+1]
		if exists {
			fmt.Println(color + "sdfs" + Reset)
		} else {
			fmt.Println("Invalid index:", index+1)
		}
	}
}
