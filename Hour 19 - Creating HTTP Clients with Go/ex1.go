//create an HTTP client that makes a GET request to http://google.com/404. Print the response code. It is 404?

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	client:= &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "http://google.com/404", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)

}