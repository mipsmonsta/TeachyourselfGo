package main

import (
	"fmt"
)

func main(){
	var cheeses [3]string //arrays are fixed sized
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourgogne"
	cheeses[2] = "Blue Cheese"
	for _, cheese := range cheeses {
		fmt.Println(cheese)
	}

}