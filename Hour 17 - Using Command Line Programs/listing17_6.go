//Using Subcommands in Command-Line Programs

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func flagUsage(){
	usageText := `Listing17_6 is an example CLI tool demonstrating sub-commands
	Usage: Listing17_6 command [arguments]
	The commands are:
	uppercase	uppercase a string
	lowercase 	lowercase a string
	Use "listing17_6 [command] --help" for more information about a command
	`
	fmt.Fprintf(os.Stderr, "%s\n", usageText)

}

func main(){
	flag.Usage = flagUsage
	uppercaseCmd:= flag.NewFlagSet("uppercase", flag.ExitOnError)
	lowercaseCmd:= flag.NewFlagSet("lowercase", flag.ExitOnError)

	if len(os.Args) == 1 { //no subcommands, so exit
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "uppercase":
		s:= uppercaseCmd.String("s", "", "Specify a string to be uppercased.")
		uppercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToUpper(*s))
	case "lowercase":
		s:= lowercaseCmd.String("s", "", "Specify a string to be lowercased")
		lowercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToLower(*s))
	default:
		flag.Usage()
	}
}