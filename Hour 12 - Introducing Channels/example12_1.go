package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
	c <- "hello"
}

func main() {
	c := make(chan string)
	go slowFunc(c)
	fmt.Println("Above is shown")
	msg:= <-c //this receive of message from channel will block
	print(msg)

}