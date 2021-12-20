package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace("    I don't want spaces around   "))
	fmt.Println(strings.TrimPrefix("McDonald", "Mc"))
}