// clean.go

package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

// Clean function should not take any arguments
func Clean() {
	var SelectMenu string
	fmt.Print("Do you want to clean Trash folder?")
	fmt.Scan(&SelectMenu)
	if SelectMenu == "n" || SelectMenu == "no" {
		os.Exit(0)
	} else if SelectMenu == "y" || SelectMenu == "yes" {
		user, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}

		username := user.Username

		RemoveErr := os.RemoveAll("/home/" + username + "/.local/share/Trash/")
		if RemoveErr != nil {
			log.Fatal(RemoveErr)
		}

		fmt.Println("Trash folder cleaned successfully.")
	}
}
