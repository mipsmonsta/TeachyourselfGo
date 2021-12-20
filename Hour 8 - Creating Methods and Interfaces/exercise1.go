//Write an interface for a taxi. You can include any methods you like,
//but you may think of things like whether the taxi is hired, how many passengers there are
//and whether the taxi is available for hire
package main

import "fmt"

type Taxi interface {
	IsHired() bool
	PassengerCapacity() int
	CanBookOnline() bool
	Company() string
}

type ComfortTaxi struct {
	companyName string
	capacity int
	isHired bool
	canBookOnline bool
}

func (t *ComfortTaxi) IsHired() bool {
	return t.isHired;
}

func (t *ComfortTaxi) PassengerCapacity() int {
	return t.capacity
}

func (t *ComfortTaxi) CanBookOnline() bool {
	return t.canBookOnline
}
func (t *ComfortTaxi) Company() string {
	return t.companyName 
}

func printTaxiInfo(d Taxi) {
	fmt.Printf("Taxi Company: %s, cab capacity %d, can book online %t, is hired %t", d.Company(), d.PassengerCapacity(), d.CanBookOnline(), d.IsHired())
}

func main(){
	c := ComfortTaxi {
		companyName: "ComfortDelgro",
		capacity: 4,
		isHired: false,
		canBookOnline: true,
	}
	printTaxiInfo(&c)
}
