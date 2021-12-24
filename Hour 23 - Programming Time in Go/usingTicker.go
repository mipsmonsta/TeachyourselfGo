package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(5 * time.Second)
	for t:= range c {
		fmt.Printf("The tick time is now % v\n", t)
	}
}