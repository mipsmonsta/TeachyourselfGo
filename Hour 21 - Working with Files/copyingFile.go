package main

import (
	"io"
	"log"
	"os"
)

func main() {

	to, err := os.OpenFile("example04.txt", os.O_RDWR | os.O_CREATE, 0644)
	if err != nil{
		log.Fatal(err)
	}

	defer to.Close()

	from, err := os.Open("example01.txt")
	if err != nil{
		log.Fatal(err)
	}

	defer from.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}