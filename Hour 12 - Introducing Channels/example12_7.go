package main

import (
	"fmt"
	"time"
)

func senderRepeat(c chan string) {
	t := time.NewTicker(time.Second * 1)
	for {
		c <- "Sending message"
		<- t.C
	}
}

func main(){
	stopChannel := make(chan string)
	c := make(chan string)

	go senderRepeat(c)
	go func() { //anonymous function simulating the stop channel to terminate the program
		time.Sleep(time.Second * 5)
		stopChannel <- "stop"
	}()
	
	for {
		select {
		case <-stopChannel: //using a stop channel to terminate the program
			return 
		case msg := <-c: 
			fmt.Println(msg)
		}
	}

}