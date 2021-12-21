//debugging HTTP Requests

//re-using listing _3 and modifying it for debugging of request and response
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main(){
	debug := flag.Bool("d", false, "Turn Debug on to dump client requests and received response")
	flag.Parse()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://ifconfig.co", nil)
	if err != nil {
		log.Fatal(err)
	}
	//NewRequest allows to customise header
	request.Header.Add("Accept", "application/json")

	//debug request
	if *debug == true {
		dumpedRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", dumpedRequest)
	}

	response, err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	
	//debug response
	if *debug == true {
		dumpedResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", dumpedResponse)
	}
	
	body, err := io.ReadAll(response.Body)
		if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)
}