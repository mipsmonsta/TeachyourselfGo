//How structs can be copied? Idea here is that by assignment, the struct will be copied over to the new variable

package main

import (
	"fmt"
)

type Phone struct {
	Model string
	HasCamera bool
	MemoryCap int
}

func main(){
	phone1 := Phone{
		Model: "Xiaomi 11T",
		HasCamera: true,
		MemoryCap: 64,
	}
	phone2 := phone1 //i.e. pass by copy
	fmt.Printf("Phone1's address: %p vs Phone2's address: %p \n", &phone1, &phone2)
}