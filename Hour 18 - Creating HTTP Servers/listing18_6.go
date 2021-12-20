//handling data from different type of requests using as GET and POST
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responsehandler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		var buffer bytes.Buffer
		for k,v := range r.URL.Query(){
			s := fmt.Sprintf("%s: %s\n", k, v[0])
			buffer.WriteString(s)
		}
		w.Write(buffer.Bytes())

	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		s := fmt.Sprintf("received: %s\n", reqBody)
		w.Write([]byte(s))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}	
	
}

func main() {
	http.HandleFunc("/", responsehandler)
	http.ListenAndServe(":8000", nil)
}
// http://localhost:8000/?foo=11&bar=22
// curl -is -X POST -d "demonstration" http://localhost:8000
