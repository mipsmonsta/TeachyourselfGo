//Using a Buffer for Concatenation
//More efficient than using +=

package main

import (
	"bytes"
	"fmt"
)

func main(){
	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}

	fmt.Println(buffer.String())

}