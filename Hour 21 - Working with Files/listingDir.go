package main

import (
	"fmt"
	"os"
)

func main() {
	entries, _ := os.ReadDir(".")
	
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("Dir name: %s\n", entry.Name())
		} else {
			fmt.Printf("File name: %s\n", entry.Name())
		}
	}

}