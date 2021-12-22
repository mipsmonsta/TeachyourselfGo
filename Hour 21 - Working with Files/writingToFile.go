package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "The fox jumps over the moon. And all was peaceful!"
	err := ioutil.WriteFile("./example03.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}