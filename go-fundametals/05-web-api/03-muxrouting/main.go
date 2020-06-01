package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerHome)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
