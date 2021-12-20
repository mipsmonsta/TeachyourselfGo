//Responding with Different Content Type

package main

import "net/http"

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
                http.NotFound(w, r)
                return
        }

        switch r.Header.Get("Accept") {
        case "application/json":
                w.Header().Set("Content-Type", "application/json; charset=utf-8")       
                w.Write([]byte(`{"message": "Hello World"}`))
        case "application/xml":
                w.Header().Set("Content-Type", "application/xml; charset=utf-8")        
                w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><Message>Hello World</Message>`))

        default:
                w.Header().Set("Content-Type", "text/plain; charset=utf-8")
                w.Write([]byte("Hello World\n"))

        }
}

func main(){
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8000", nil)
}

//use curl to try out the different content-type
//url -is -H 'Accept: application/xml' http:/localhost:8000/
