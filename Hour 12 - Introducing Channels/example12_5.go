package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string){
	t:= time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<- t.C //this will block, waiting for ticker to tick
		
	}
}

func main() {
	messages := make(chan string)
	go slowFunc(messages)
	for {
		msg := <- messages //this will block, waiting for ticker to tick
		fmt.Println(msg)
	}
}