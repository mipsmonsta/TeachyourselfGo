package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos"
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Accept", "application/json")

	//make the request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var todos []todo
	err = json.NewDecoder(response.Body).Decode(&todos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", todos)

}