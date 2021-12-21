package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

func main() {
	response, err := http.Get("http://ifconfig.co/")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//get the external facing IP address in the WAN
	fmt.Println("Body:", string(body)) //if no string casting a slide of uint8 will be printed
	fmt.Println(reflect.TypeOf(body))

}