package main

import (
	"fmt"
	"time"
)

func ping(c chan string, d time.Duration, name string){
	time.Sleep(d);
	c <- fmt.Sprint("ping received with %s", name)
}

func main(){
	channel1 := make(chan string)
	channel2 := make(chan string)
	go ping(channel1, time.Second * 4, "first channel")
	go ping(channel2, time.Second * 5, "second channel")
	
	select {
	case msg := <- channel1:
		fmt.Println("received", msg)
	case msg2 := <- channel2:
		fmt.Println("received", msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("no message received")
	}
}