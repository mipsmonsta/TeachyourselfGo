//Using panic to halt execution

package main

import "fmt"

func main(){
	fmt.Println("Some execution")
	panic("Something in execution went wrong and we shan't continue")
	fmt.Println("This will not be printed")
}