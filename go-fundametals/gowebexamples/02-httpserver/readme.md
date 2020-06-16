# Introduction
In this example you will learn how to create a basic HTTP server in Go. First, let’s talk about what our HTTP server should be capable of. A basic HTTP server has a few key jobs to take care of.

Process dynamic requests: Process incoming requests from users who browse the website, log into their accounts or post images.
Serve static assets: Serve JavaScript, CSS and images to browsers to create a dynamic experience for the user.
Accept connections: The HTTP Server must listen on a specific port to be able to accept connections from the internet.