//Using the Robot Interface

package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
	Talk() string
}

type T850 struct {
	Name string
}


func (a *T850) PowerOn() error {
	return nil
}

func (a *T850) Talk() string {
	return fmt.Sprintf("I am %s", a.Name)
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func (a *R2D2) Talk() string {
	return "I am R2D2!"
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func Speak(r Robot) string {
	return r.Talk()
}

func (r Robot) SpeakAgain() string { //wrong since you are associate a concrete method to an interface, which is not a concrete type
	return r.SpeakAgain() 
}

func main(){
	t:= T850{
		Name: "The Terminator",
	}
	r:= R2D2{
		Broken: true,
	}
	err := Boot(&r) //interface is not a concrete type and is a pointer, so you need to use &r
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	fmt.Println(Speak(&t))
	fmt.Println(Speak(&r))

}