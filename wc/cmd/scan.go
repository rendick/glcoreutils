package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Scan(OutputText string) (int, int) {
	// lines
	file_count, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file_count.Close()

	scanner := bufio.NewScanner(file_count)

	var countLines int
	for scanner.Scan() {
		countLines++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return 0, 0
	}

	file_count.Seek(0, 0)

	// words
	var countWords = make(map[string]int)
	scanner = bufio.NewScanner(file_count)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			countWords[word]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return 0, 0
	}

	// size
	file_size, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(OutputText, countLines, len(countWords), file_size.Size(), os.Args[1])

	return countLines, len(countWords)
}
