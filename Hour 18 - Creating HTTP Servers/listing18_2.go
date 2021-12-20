//setting a 404 Not found response

package main

import (
	"fmt"
	"net/http"
)

func helloworldHandler(w http.ResponseWriter, r *http.Request){

	fmt.Println(r.URL)
	//fmt.Println(r.Host)
	if r.URL.Path != "/" { // "/" will handle all paths, so need to ensure
						   // handler's behaviour is to block
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello World\n"))
}

func main() {
	http.HandleFunc("/", helloworldHandler)
	http.ListenAndServe(":8000", nil)
}