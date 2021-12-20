package main

import "fmt"

func main() {
	fourSlice := make([]string, 4)                               // the capacity is a guidance for slice
	fourSlice = append(fourSlice, "one", "two", "three", "four") //append is a variadic function, can take in more than 1 arguments
	// for _, value := range fourSlice {
	// 	fmt.Println(value)
	// }
	//create new slice and copy only the third and fourth element into it
	var newSlice = make([]string, 2) //shorform is newSlice := make...
	//newSlice = append(newSlice, fourSlice[2:]...)
	newSlice = append(newSlice, "five", "six") //must have elements in it, ease copying in next line will not work i.e. only empty newSlice
	copy(newSlice, fourSlice[2:])
	for _, value := range newSlice {
		fmt.Println(value)
	}
	fmt.Println(newSlice)

	newNewSlice := []string{"five", "six"}

	for _, value := range newNewSlice {
		fmt.Println(value)
	}
	



}