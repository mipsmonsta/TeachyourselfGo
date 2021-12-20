//checking type of structs
package main

import (
	"fmt"
	"reflect"
)

type Piano struct {
	Model string
	color string
}

func main(){
	piano1 := Piano {
		Model: "Yamaha",
		color: "blue", //private to the package
	}

	fmt.Printf("Piano: %+v \n", piano1)
	fmt.Println(reflect.TypeOf(piano1))

}
