//Responding to different types of Requests

package main

import "net/http"

func respondingHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		w.Write([]byte("Received a GET request\n"))
		
	case "POST":
		w.Write([]byte("Received a POST request\n"))

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))

	}
}

func main(){
	http.HandleFunc("/", respondingHandler)
	http.ListenAndServe(":8000", nil)
}

// curl -is -X POST http://localhost:8000
// curl -is -X GET http://localhost:8000
// curl -is -X DELETE http://localhost:8000

