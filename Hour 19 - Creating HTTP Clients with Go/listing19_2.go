//post request to https://httpbin.org/post

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main(){
	postData := strings.NewReader(`{
		"some": "json"
	}`)
	response, err := http.Post("https://httpbin.org/post", "application/json", postData)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}