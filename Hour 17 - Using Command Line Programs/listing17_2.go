//parsing command-line flags

package main

import (
	"flag"
	"fmt"
)

func main(){
	s := flag.String("s", "Default String value", "Set the string value for 's' flag")
	flag.Parse()
	fmt.Println("Value of s", *s)
}