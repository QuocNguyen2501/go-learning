# Middleware (Basic)
This example will show how to create basic logging middleware in Go.

A middleware simply takes a http.HandlerFunc as one of its parameters, wraps it and returns a new http.HandlerFunc for the server to call.