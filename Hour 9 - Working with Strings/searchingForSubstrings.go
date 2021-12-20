package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("surface", "ace")) //4
	fmt.Println(strings.Index("cannot", "find")) //-1
}