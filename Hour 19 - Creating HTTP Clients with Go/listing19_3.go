//using a custom client - should use client for further control of HTTP Requests (e.g. set headers)
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://ifconfig.co", nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
		if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}