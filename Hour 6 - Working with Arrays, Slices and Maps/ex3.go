package main

import (
	"fmt"
)

func main(){
	elements := make(map[string]string)
	elements["p"] = "Paragraph"
	elements["img"] = "Image"
	elements["h1"] = "Heading One"
	elements["h2"] = "Heading Two"

	//using map literals
	eles := map[string]string{"p": "Paragraph", "img": "Image", "h1": "Heading One", "h2": "Heading Two"}
	for key, ele := range eles {
		result := fmt.Sprintf("key: %s, value: %s", key, ele)
		fmt.Println(result)
	}
}