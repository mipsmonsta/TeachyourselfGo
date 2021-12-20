package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	t := time.NewTicker(time.Second * 2)
	for {
		c <- "ping"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	go slowFunc(messages)
	msg := <- messages
	fmt.Println(msg)
}