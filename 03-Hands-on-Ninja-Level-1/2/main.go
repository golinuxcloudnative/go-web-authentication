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
	http.HandleFunc("/decode", decode())
	http.ListenAndServe(":8000", nil)

}

func decode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			errStr := fmt.Sprintf("%s method not allowed", r.Method)
			http.Error(w, errStr, http.StatusMethodNotAllowed)
			return
		}
		people := []Person{}
		err := json.NewDecoder(r.Body).Decode(&people)
		if err != nil {
			log.Printf("could not decode data, %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Data decoded - %v", people)
		w.WriteHeader(200)
		w.Write([]byte("successfully decoded"))
	}
}
