package main

import (
	"fmt"
)

func main() {
	s:= "Hello"
	fmt.Printf("%q \n", s[1]) //need to use %q flag
	fmt.Printf("%b \n", s[1]) //show binary representation
}