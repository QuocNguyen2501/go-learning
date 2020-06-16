# Introduction
Go’s net/http package provides a lot of functionalities for the HTTP protocol. One thing it doesn’t do very well is complex request routing like segmenting a request url into single parameters. Fortunately there is a very popular package for this, which is well known for the good code quality in the Go community. In this example you will see how to use the gorilla/mux package to create routes with named parameters, GET/POST handlers and domain restrictions.

go get -u github.com/gorilla/mux