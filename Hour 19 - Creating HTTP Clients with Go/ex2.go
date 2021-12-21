//create an HTTP client that makes a POST request to https://httpbin.org/post. Post some data and examine the reponse to see that it was correctly posted.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main(){
	url := "https://httpbin.org/post"

	client := &http.Client{}
	reader := strings.NewReader(`
		<?xml version = "1.0" encoding = "UTF-8" ?>
		<some>
		xml string
		</some>
	`)
	request, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/xml; utf-8")
	
	dumpedRequest, err:= httputil.DumpRequest(request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", dumpedRequest)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}



	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	
}