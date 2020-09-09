package main

import (
	"fmt"
	"net/http"
)

func sayHelloworld(w http.ResponseWriter, r *http.Request) { // <-- Handler
	fmt.Fprintf(w, "Hello world")
}

func doSomething(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do something")
}

func main() {
	http.HandleFunc("/hello-world", sayHelloworld)
	http.HandleFunc("/do-something", doSomething)
	http.ListenAndServe(":8080", nil)
}
