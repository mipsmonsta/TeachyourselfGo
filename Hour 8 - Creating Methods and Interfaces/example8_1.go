//declaring and calling a method i.e. not a function

package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name string
	Rating float64
}

//declaring the method
func (m *Movie) Summary() string { //notice the use of *movie to indicate pass by reference of movie
	r := strconv.FormatFloat(m.Rating, 'f', 2, 64)
	return m.Name + ", " + r
}

func main(){
	m := Movie{
		Name: "Titanic",
		Rating: 9.612,
	}
	fmt.Printf("This is about %s", m.Summary())
}