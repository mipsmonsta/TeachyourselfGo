package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//read in as []byte
	fileBytes, err := ioutil.ReadFile("example01.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(fileBytes))
}