package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string   `json:"name"` //tag, note that no spacing allowed in tag, else won't work
	Age     int      `json:"age"`                // `json: "age" is a tag and here is used to ensure json are in camelCase rather than in caps that is common for GoLang`
	Hobbies []string `json:"hobbies,omitempty"` //using slice here, omitempty here ensures that zero values are omitted when marshalling
}

func main() {
	h := []string{"Badminton", "Gaming", "Toast-Mastering", "Wood works"}
	p := Person{
		Name:    "John",
		Age:     45,
		Hobbies: h,
	}

	jsonByteData, err := json.Marshal(p) // struct is an interface, so can be accepted
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)

	fmt.Println(jsonStringData)

	//decoding json
	aJsonByteData := []byte(jsonStringData)
	anotherP := Person{}
	err = json.Unmarshal(aJsonByteData, &anotherP)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(anotherP)
}