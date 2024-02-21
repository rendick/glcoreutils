package cmd

import (
	"log"
	"os"
	"strconv"
	"time"
)

func Sleep() {
	conv2int, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Duration(conv2int) * time.Second)
}
