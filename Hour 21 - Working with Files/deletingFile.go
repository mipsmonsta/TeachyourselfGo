package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("./folderToBeFound")
	if err != nil {
		log.Fatal(err)
	}
}