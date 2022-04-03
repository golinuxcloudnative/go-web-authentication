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
	http.HandleFunc("/decode", decode())
	http.ListenAndServe(":8000", nil)

}

func encode(p []Person) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errStr := fmt.Sprintf("%s method not allowed", r.Method)
			http.Error(w, errStr, http.StatusMethodNotAllowed)
			return
		}
		name := r.URL.Query().Get("name")
		for _, person := range p {
			if person.FirstName == name {
				err := json.NewEncoder(w).Encode(person)
				if err != nil {
					log.Printf("could not encode data, %v", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}
		http.Error(w, "could not find the user", http.StatusNotFound)
	}
}

func decode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			errStr := fmt.Sprintf("%s method not allowed", r.Method)
			http.Error(w, errStr, http.StatusMethodNotAllowed)
			return
		}
		var p1 Person
		err := json.NewDecoder(r.Body).Decode(&p1)
		if err != nil {
			log.Printf("could not decode data, %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("Data decoded - %v", p1)
		w.WriteHeader(200)
		w.Write([]byte("successfully decoded"))
	}
}
