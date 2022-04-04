package main

import (
	"log"
	"net/http"
)

var (
	user   = "admin"
	passwd = "admin"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestUser, requestPasswd, ok := r.BasicAuth()
		if !ok {
			log.Printf("request forbidden, no authentication")
			w.Header().Add("WWW-Authenticate", "Basic realm=\"Access to secret site\"")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if requestUser != user || requestPasswd != passwd {
			log.Printf("wrong user or password, user: %s", requestUser)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Printf("user successfully logged in, %s", requestUser)
		w.Write([]byte("<h1>User authenticated</h1>"))
		return
	})

	http.ListenAndServe(":8000", nil)

}
