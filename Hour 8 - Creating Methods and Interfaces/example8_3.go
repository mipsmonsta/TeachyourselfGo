//passing by pointer reference to a method

package main

import (
	"fmt"
	"strconv"
)

type Triangle struct {
	base float64
	height float64
}

func (t *Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t *Triangle) Area2DP() string {
	areaf := t.area()
	return strconv.FormatFloat(areaf, 'f', 2, 64)
}

func main(){
	t := Triangle{
		base: 3.2,
		height: 5.3,
	}
	fmt.Printf("Area of Triange %+v is %f \n and %s in 2 d.p.", t, t.area(), t.Area2DP())
}