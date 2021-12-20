// Demonstrating the different value types accepted by flag

package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("s", "A default string value", "String help text")
	i := flag.Int("i", 42, "Meaning of life")
	b := flag.Bool("b", false, "Bool help text")
	flag.Parse()

	fmt.Println("value of s:", *s)
	fmt.Println("Value of i:", *i)

	//note that for bool flag, simply setting it in command line will always give true
	//only when -b flag is omitted default value above of false will show
	fmt.Println("Value of b:", *b)
	fmt.Printf("Value of b again: %t\n", *b)
}