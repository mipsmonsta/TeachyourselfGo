//Comparing Structs with the Same Values Assigned to Data Elements
package main

import (
	"fmt"
)

type Drink struct {
	Name string
	Ice bool
}

func main(){
	drink1 := Drink{
		Name: "Bandung",
		Ice: true,
	}

	drink2 := Drink{
		Name: "Rose Syrup Milk",
		Ice: true,
	}

	drink3 := Drink{
		Name: "Bandung",
		Ice: true,
	}

	if (drink3 == drink1){
		fmt.Printf("Drink 1: %v is equivalent to Drink3: %v \n", drink1, drink3)
	}

	if (drink1 != drink2){
		fmt.Printf("Drink 1: %v is not equivalent to Drink2: %v \n", drink1, drink2)
	}

}