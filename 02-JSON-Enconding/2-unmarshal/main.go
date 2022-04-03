package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string
}

func main() {

	data := `[
		{
		  "FirstName": "Luiz Henrique"
		},
		{
		  "FirstName": "Amora"
		}
	  ]`

	people := []Person{}

	err := json.Unmarshal([]byte(data), &people)

	if err != nil {
		log.Fatalf("cannot unmarshal, %v", err)
	}

	fmt.Println(people)
}
