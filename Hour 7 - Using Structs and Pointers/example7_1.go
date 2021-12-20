//declaring and initializing a struct
package main

import (
	"fmt"
)

type Movie struct {
	Name string //no commas after
	Rating float32 //when you use uppercase, it's a public assessible member of the struct
}

func main(){
	//using short hand variable to initialize a struct
	m := Movie{ //"squiggly brackets" 
		Name: "Predator",
		Rating: 10, //the last commas is required
	}

	fmt.Printf("%+v \n", m) //display keys and values
	fmt.Printf("%v \n", m) //only display values

	//use new to initialize the struct
	m1 := new(Movie)
	m1.Name = "Matrix"
	m1.Rating = 8

	fmt.Printf("%+v \n", m1)
}	