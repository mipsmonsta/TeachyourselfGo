package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("!asdf$22$2"))
	fmt.Println(re.MatchString("roger"))
}