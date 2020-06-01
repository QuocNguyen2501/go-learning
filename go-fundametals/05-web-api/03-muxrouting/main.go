package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomeHandler")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductsHandler")
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ArticlesHandler")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/{key}", productsHandler)
	r.HandleFunc("/articles/{id:[0-9]+}", articlesHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
