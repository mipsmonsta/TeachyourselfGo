//Errors can be created using the errors package
//and also using fmt 

package main

import (
	"errors"
	"fmt"
)

func main(){
	//first creation ways
	err := errors.New("Created error. Something went wrong")
	if err != nil {
		fmt.Println(err)
	}

	var code = "Err2213"

	//second creation way
	err2 := fmt.Errorf("Error code: %s. Something went wrong", code)
	if err2 != nil {
		fmt.Println(err2) 
	}
	
}