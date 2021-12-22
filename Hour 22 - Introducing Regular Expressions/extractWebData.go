package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {

	client := &http.Client{}

	url := "http://petition.parliament.uk/petitions"

	request, err:= http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	src := string(body)

	//fmt.Println(src)

	re:= regexp.MustCompile("<h2>.*</h2>")
	headers := re.FindAllString(src, -1)
	fmt.Println("Before cleaning:")
	for _, header := range headers {
		fmt.Println(header)
	}

	fmt.Println("After cleaning:")
	an:= regexp.MustCompile("<[^<]*>")
	pairs := map[string]string {
		"&#34;":"\"",
		"&amp;":"&",
		"&#39;":"'",
		"&#43;":"+",
		"&lt;":"<",
		"&gt;":">",
	}

	// 0:    "\uFFFD",
	// '"':  "&#34;",
	// '&':  "&amp;",
	// '\'': "&#39;",
	// '+':  "&#43;",
	// '<':  "&lt;",
	// '>':  "&gt;",
	for _, header := range headers {
		header = an.ReplaceAllString(header, "") //remove <>
		for k,v := range pairs{
			en := regexp.MustCompile(k)
			header = en.ReplaceAllString(header, v)
		}
		
		fmt.Println(header)
	}

}