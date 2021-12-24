package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s := "2011-01-02T19:00:00+08:00"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The hour is %v\n", t.Hour())
	fmt.Printf("UNIX time is %v\n", t.Unix()) //in seconds 
	fmt.Printf("UNIX time ms is %v\n", t.UnixMicro())
	fmt.Printf("The day of week is %v\n", t.Weekday())
 }