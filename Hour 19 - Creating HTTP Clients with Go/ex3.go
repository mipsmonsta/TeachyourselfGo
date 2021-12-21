//create an HTTP client that makes a GET request to https://httpbin.org/user-agent. Modify the 'User-Agent' header to be "GolangBot".
//make a request and examine the user-agent value in the response. If it says "GolangBot", you just set the User Agent!

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	url := "https://httpbin.org/user-agent"
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("User-Agent", "GolangBot")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Header.Get("Content-Type"))
	
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}