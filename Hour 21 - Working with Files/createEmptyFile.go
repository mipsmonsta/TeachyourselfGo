package main

import (
	"io/ioutil"
	"log"
)

func main() {

	b := make([]byte, 0) //"empty bytes"
	err := ioutil.WriteFile("example02.txt", b, 0644) //-rw-rw----
	if err != nil {
		log.Fatal(err)
	}
}