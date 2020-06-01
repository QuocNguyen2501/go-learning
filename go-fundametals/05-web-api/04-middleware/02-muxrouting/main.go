package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerHome).Use(loggingMiddleware)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
