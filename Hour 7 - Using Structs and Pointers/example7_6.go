//setting Default Values Usinga Constructor Function
package main

import (
	"fmt"
)

type Alarm struct {
	Time string
	Sound bool
}

func NewAlarm(time string) Alarm { //Constructor Function
	a := Alarm{
		Time: time,
		Sound: true,
	}
	return a
}

func main(){
	fmt.Printf("%+v\n", NewAlarm("23:00"))
}