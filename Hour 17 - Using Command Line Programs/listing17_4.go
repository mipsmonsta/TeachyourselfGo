//Customising usage text of flag

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	command := os.Args[0]
	_, filePath := filepath.Split(command)
	flag.Usage = func() { //custom usage function
		usageText := `Usage %s [OPTION]
		An example of customising usage output
		
		-s, --s		example string argument, default: String help text
		-i, --i		example integer argument, default: Int help text
		-b, --b		example boolean argument, default: Bool help text`
		
		usageText = fmt.Sprintf(usageText, filePath)
		
		fmt.Fprintf(os.Stderr, "%s\n", usageText)

	}

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