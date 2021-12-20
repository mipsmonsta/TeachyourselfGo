//Returning an error from a Function

package main

import "fmt"

func Half(numberToHalf int) (int, error){
	
	if numberToHalf % 2 != 0 {
		return -1, fmt.Errorf("Cannot half number: %d", numberToHalf)
	}

	return numberToHalf / 2, nil
	
}

func main(){
	n, err := Half(19)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("Halved number result is %d \n", n)

}