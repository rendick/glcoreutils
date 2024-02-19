package main

import (
	"fmt"
	"runtime"

	"github.com/rendick/glcoreutils/internal/rm/cmd"
)

func main() {
	if runtime.GOOS == "linux" {
		cmd.Switch()
	} else {
		fmt.Println("You aren't running Linux!")
	}
}
