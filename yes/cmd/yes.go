package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	Color1 = "\033[31m"
	Color2 = "\033[32m"
	Color3 = "\033[33m"
	Color4 = "\033[35m"

	Reset = "\033[0m"
)

func Yes() {
	color_arr := map[int]string{
		1: Color1,
		2: Color2,
		3: Color3,
		4: Color4,
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
