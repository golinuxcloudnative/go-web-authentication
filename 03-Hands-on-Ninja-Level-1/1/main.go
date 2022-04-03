package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	FirstName string
}

func main() {

	people := []Person{
		{FirstName: "Luiz Henrique"},
		{FirstName: "Amora"},
		{FirstName: "Rafael"},
		{FirstName: "Daniel"},
		{FirstName: "Todd"},
	}

	http.HandleFunc("/encode", encode(people))
	http.ListenAndServe(":8000", nil)

}

func encode(p []Person) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errStr := fmt.Sprintf("%s method not allowed", r.Method)
			http.Error(w, errStr, http.StatusMethodNotAllowed)
			return
		}
		err := json.NewEncoder(w).Encode(p)
		if err != nil {
			log.Printf("could not encode data, %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
