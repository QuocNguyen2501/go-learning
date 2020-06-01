package main

import (
	"fmt"
	"log"
	"net/http"
)

func loggingMiddleware(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		fn(w, r, name)
	}
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func main() {
	http.HandleFunc("/home", loggingMiddleware(handlerHome))
	http.ListenAndServe(":8080", nil)
}
