//Error Handling when reading a file
//In Go error is a type and can be returned
//There is no try-catch-finally.
//Error is also not an exception i.e. you as the programmer
//decides to act when receiving and error. This is unlike
//other programming langauge when an error will crash the program.

package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	file, err := ioutil.ReadFile("foo.txt") //err is returned
	if err != nil {
		fmt.Println(err) //you decide how to handle
	} 
	fmt.Printf("%s", file)
}