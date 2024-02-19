package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// sudo chmod 777 ~/.local/share/Trash

func Deletion() {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username

	dir := flag.String("r", "", "folder")
	flag.Parse()

	if *dir == "" {
		log.Fatal("Please provide a folder to delete using the -r flag.")
	}

	absDir, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}

	trashPath := "/home/" + username + "/.local/share/Trash/"
	trashDir := filepath.Join(trashPath, filepath.Base(absDir))

	_, err = os.Stat(absDir)
	if os.IsNotExist(err) {
		log.Fatalf("Directory %s does not exist.", absDir)
	}

	_, err = os.Stat(trashPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(trashPath, 0750)
		if err != nil {
			log.Panic(err)
		}
	}

	err = os.Rename(absDir, trashDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Moved %s to Trash: %s\n", absDir, trashDir)
}

func main() {
	Deletion()
}
