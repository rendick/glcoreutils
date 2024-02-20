package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	Color1 = "\033[31m"
	Color2 = "\033[32m"
	Color3 = "\033[33m"
	Color4 = "\033[35m"

	Reset = "\033[0m"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <filename>")
		os.Exit(1)
	}

	color_arr := map[int]string{
		1: Color1,
		2: Color2,
		3: Color3,
		4: Color4,
	}

	file_count, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(file))

	rand.Seed(time.Now().UnixNano())
	gen := rand.Perm(4)
	rand.Shuffle(len(gen), func(i, j int) { gen[i], gen[j] = gen[j], gen[i] })

	for _, index := range gen {
		color := color_arr[index+1]
		if color != "" {
			fmt.Print(color + strings.TrimSpace(string(file)) + Reset)
		}
	}

	scanner := bufio.NewScanner(file_count)
	scanner.Split(bufio.ScanLines)

	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(Color1+"Number of lines in file: "+Reset, count)
	}
}
