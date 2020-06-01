package main

import (
	"fmt"
	"net/http"
)

func handlerName(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[7:]
		fn(w, r, name)
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request, name string) {
	salutation := getSalutation(name)
	fmt.Fprintf(w, salutation+name)
}

func getSalutation(name string) string {
	switch name {
	case "Bob":
		return "Mr "
	case "Mary":
		return "Mrs "
	default:
		return "Due "
	}
}

func main() {
	http.HandleFunc("/hello/", handlerName(handlerHello))
	http.ListenAndServe(":8080", nil)
}
