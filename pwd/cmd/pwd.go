package cmd

import (
	"fmt"
	"math/rand"
	"os"
)

var (
	Color1 = "\033[31m"
	Color2 = "\033[32m"
	Color3 = "\033[33m"
	Color4 = "\033[35m"
	Color5 = "\033[37m"

	Reset = "\033[0m"
)

func GetDir() {
	color_arr := map[int]string{
		1: Color1,
		2: Color2,
		3: Color3,
		4: Color4,
		5: Color5,
	}
	gen := rand.Perm(5)
	for _, index := range gen {
		color := color_arr[index+1]

		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		} else {
			fmt.Println(color + pwd + Reset)
			os.Exit(0)
		}
	}
}
