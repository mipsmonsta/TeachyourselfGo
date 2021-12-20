//Converting Types to a String

package main

import (
	"fmt"
	"strconv"
)

func main(){
	var i int = 1

	iStr := strconv.Itoa(i) //convert int to string
	resultStr := iStr + " egg"
	fmt.Println(resultStr)

}