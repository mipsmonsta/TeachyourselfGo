//Must build the code before code will work

package main

import (
	"fmt"
	"os"
)

//Main will print out the arguments
func main() {
	for i, arg := range os.Args {
		fmt.Println("argument", i, "is", arg)
	}
}