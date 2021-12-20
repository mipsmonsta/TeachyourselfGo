//demonstrating nested structs
package main

import "fmt"

type Superhero struct {
	Name string
	Age int
	Address Address
}

type Address struct {
	Number int
	Street string
	City string
}

func main() {
	e := Superhero{
		Name: "Batman",
		Age: 32,
		Address: Address{
			Number: 27,
			Street: "Gotham City Apartment",
			City: "Gotham City",
		},
	}
	fmt.Printf("Superhero -> %+v", e)
}