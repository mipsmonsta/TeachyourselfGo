package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s1 := "2021-12-25T12:43:00+08:00"
	t, err := time.Parse(time.RFC3339, s1)
	if err != nil {
		log.Fatal(err)
	}
	nt := t.Add(2 * time.Hour)
	fmt.Println(nt)

	nt = t.Add(-60 * time.Second) //subtract
	fmt.Println(nt)

	fmt.Println(nt.Before(t))
}